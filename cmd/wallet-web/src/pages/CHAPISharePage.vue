<!--
 * Copyright SecureKey Technologies Inc. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
-->

<script setup>
import { ref, inject, computed, onMounted, markRaw, toRaw } from 'vue';
import { useStore } from 'vuex';
import {
  filterCredentialsByType,
  getCredentialIcon,
  WalletGetByQuery,
  prepareCredentialManifest,
  resolveManifest,
} from '@/mixins';
import SpinnerIcon from '@/components/icons/SpinnerIcon.vue';
import { useI18n } from 'vue-i18n';
import { CredentialManager } from '@trustbloc/wallet-sdk';

// Local Variables
const errors = ref([]);
const requestOrigin = ref('');
const loading = ref(true);
const sharing = ref(false);
const credsFound = ref([]);
const issuersFound = ref([]);
const manifestCredType = 'IssuerManifestCredential';
const wallet = ref(null);
const presentation = ref(null);

// Computed
const showCredentialsMissing = computed(() => !credsFound.value.length);

// Hooks
const store = useStore();
const { t } = useI18n();
const protocolHandler = inject('protocolHandler');

// Store Getters
const currentUser = computed(() => store.getters['getCurrentUser']);
const agentOpts = computed(() => store.getters['getAgentOpts']);
const getStaticAssetsUrl = computed(() => store.getters['getStaticAssetsUrl']);
const getAgentInstance = computed(() => store.getters['agent/getInstance']);
const credentialManifests = computed(() => store.getters['getCredentialManifests']);

// Methods
onMounted(async () => {
  const { user, token } = currentUser.value.profile;
  wallet.value = markRaw(
    new WalletGetByQuery(getAgentInstance.value, protocolHandler.value, agentOpts.value, user)
  );
  const credentialManager = new CredentialManager({ agent: getAgentInstance.value, user });
  try {
    // make sure mediator is connected
    await wallet.value.connectMediator();
  } catch (e) {
    console.error(e);
  }
  requestOrigin.value = protocolHandler.value.requestor();
  try {
    presentation.value = markRaw(await wallet.value.getPresentationSubmission(token));
  } catch (e) {
    errors.value.push(e);
    console.error('get credentials failed,', e);
    loading.value = false;
    return;
  }
  const credentials = toRaw(presentation.value.verifiableCredential);
  let manifest = prepareCredentialManifest(
    presentation.value,
    credentialManifests.value,
    requestOrigin.value
  );
  credsFound.value = await resolveManifest(credentialManager, credentialManifests.value, token, {
    manifest,
    response: presentation.value,
  });
  issuersFound.value = filterCredentialsByType(credentials, [manifestCredType], true);
  loading.value = false;
});

function toggleDetails(credential) {
  credential.showDetails = !credential.showDetails;
}

async function createPresentation() {
  sharing.value = true;
  try {
    await wallet.value.createAndSendPresentation(currentUser.value, presentation.value);
  } catch (e) {
    console.error(e);
    errors.value.push('share credentials failed,', e);
    sharing.value = false;
  }
  sharing.value = false;
}

function cancel() {
  wallet.value.cancel();
}

function getCredentialIconFunction(credential) {
  const cred = toRaw(credential);
  return cred?.styles?.thumbnail?.uri?.includes('https://')
    ? cred?.styles?.thumbnail?.uri
    : getCredentialIcon(getStaticAssetsUrl.value, cred?.styles?.thumbnail?.uri);
}

function focusStyleColor(color) {
  return {
    'background-color': color,
    '--focus-color': color,
  };
}
</script>

<template>
  <!-- Loading state -->
  <div v-if="loading" class="flex justify-center items-start w-screen h-screen">
    <div
      class="flex justify-center items-center w-full max-w-md h-80 bg-gray-light border-neutrals-black md:border md:border-t-0"
    >
      <SpinnerIcon />
    </div>
  </div>
  <!-- Sharing State -->
  <div v-else-if="sharing" class="flex justify-center items-start w-screen h-screen">
    <div
      class="flex flex-col justify-center items-center w-full max-w-md h-80 bg-gray-light border-neutrals-black md:border md:border-t-0"
    >
      <SpinnerIcon />
      <span class="mt-8 text-base text-neutrals-dark">{{
        t('CHAPI.Share.sharingCredential')
      }}</span>
    </div>
  </div>
  <!-- Error State -->
  <div
    v-else-if="!showCredentialsMissing && errors.length"
    class="flex justify-center items-start w-screen h-screen"
  >
    <div
      class="flex flex-col justify-center items-center w-full max-w-md h-auto bg-gray-light border-neutrals-black md:border md:border-t-0"
    >
      <div class="flex flex-col justify-start items-center py-16 px-5">
        <img src="@/assets/img/icons-error.svg" />
        <span class="mt-5 mb-3 text-xl font-bold text-center text-neutrals-dark">{{
          t('CHAPI.Share.Error.heading')
        }}</span>
        <span class="text-lg text-center text-neutrals-medium">{{
          t('CHAPI.Share.Error.body')
        }}</span>
      </div>
      <div
        class="flex flex-row justify-center items-center py-4 px-5 w-full bg-neutrals-magnolia border-t border-neutrals-thistle"
      >
        <button id="share-credentials-ok-btn" class="btn-primary" @click="cancel">
          {{ t('CHAPI.Share.Error.tryAgain') }}
        </button>
      </div>
    </div>
  </div>
  <!-- Credentials Missing State -->
  <div v-else-if="showCredentialsMissing" class="flex justify-center items-start w-screen h-screen">
    <div
      class="flex flex-col justify-center items-center w-full max-w-md h-auto bg-gray-light border-neutrals-black md:border md:border-t-0"
    >
      <div class="flex flex-col justify-start items-center py-16 px-5">
        <img src="@/assets/img/icons-error.svg" />
        <span class="mt-5 mb-3 text-xl font-bold text-center text-neutrals-dark">{{
          t('CHAPI.Share.CredentialsMissing.heading')
        }}</span>
        <span class="text-lg text-center text-neutrals-medium">{{
          t('CHAPI.Share.CredentialsMissing.body')
        }}</span>
      </div>
      <div
        class="flex flex-row justify-center items-center py-4 px-5 w-full bg-neutrals-magnolia border-t border-neutrals-thistle"
      >
        <button id="share-credentials-ok-btn" class="btn-outline" @click="cancel">
          {{ t('CHAPI.Share.CredentialsMissing.ok') }}
        </button>
      </div>
    </div>
  </div>
  <!-- Main State -->
  <div
    v-else
    class="flex overflow-scroll justify-center items-start w-screen h-screen max-h-screen"
  >
    <div class="w-full max-w-md bg-gray-light border-neutrals-black md:border md:border-t-0">
      <div class="p-5">
        <!-- Heading -->
        <div class="flex flex-row justify-start items-start mb-4 w-full">
          <div class="flex-none w-12 h-12 border-opacity-10">
            <!-- TODO: issue-1055 Read meta data from external urls -->
            <img src="@/assets/img/generic-issuer-icon.svg" />
          </div>
          <div class="flex flex-col pl-3">
            <span class="flex-1 mb-1 text-sm font-bold text-left text-neutrals-dark text-ellipsis">
              <!-- TODO: issue-1055 Read meta data from external urls -->
              Verifier
            </span>
            <div class="flex flex-row justify-center items-center">
              <img src="@/assets/img/small-lock-icon.svg" />
              <span class="flex-1 pl-1 text-xs text-left text-neutrals-medium text-ellipsis">
                {{ requestOrigin }}
              </span>
            </div>
          </div>
        </div>

        <span class="text-neutrals-dark">{{
          t('CHAPI.Share.headline', credsFound.length, { issuer: 'Verifier' })
        }}</span>
        <div v-if="credsFound.length" class="flex flex-col justify-start items-center my-6 w-full">
          <ul class="space-y-5 w-full">
            <li v-for="(credential, index) in credsFound" :key="index">
              <!-- Credential Preview -->
              <button
                :class="[
                  `group inline-flex items-center rounded-xl p-5 text-sm md:text-base font-bold border w-full h-20 md:h-24 focus-within:ring-2 focus-within:ring-offset-2 credentialPreviewContainer notWhiteCredentialPreview`,
                  credential.styles.background.color !== '#fff'
                    ? `border-neutrals-black border-opacity-10`
                    : `bg-neutrals-white border-neutrals-thistle hover:border-neutrals-chatelle`,
                ]"
                :style="focusStyleColor(credential.styles.background.color)"
                @click="toggleDetails(credential)"
              >
                <div class="flex-none w-12 h-12 border-opacity-10">
                  <img :src="getCredentialIconFunction(credential)" />
                </div>
                <div class="grow p-4">
                  <span
                    :class="[
                      `text-sm md:text-base font-bold text-left text-ellipsis`,
                      credential.styles.background.color !== '#fff'
                        ? `text-neutrals-white`
                        : `text-neutrals-dark`,
                    ]"
                  >
                    {{ credential.title }}
                  </span>
                </div>
              </button>
              <!-- Credential Details -->
              <div
                v-if="credential.showDetails"
                class="flex flex-col justify-start items-start mt-5 w-full md:mt-6"
              >
                <span class="py-3 text-base font-bold text-neutrals-dark">What's being shared</span>

                <!-- TODO: move this to reusable components -->
                <table class="w-full border-t border-neutrals-chatelle">
                  <tr
                    v-for="(property, key) of credential.properties"
                    :key="key"
                    class="border-b border-neutrals-thistle border-dotted"
                  >
                    <td class="py-4 pr-6 pl-3 text-neutrals-medium">{{ property.label }}</td>
                    <td
                      v-if="property.schema.contentMediaType != 'image/png'"
                      class="py-4 pr-6 pl-3 text-neutrals-dark break-words"
                    >
                      {{ property.value }}
                    </td>
                    <td
                      v-if="property.schema.contentMediaType === 'image/png'"
                      class="py-4 pr-6 pl-3 text-neutrals-dark break-words"
                    >
                      <img :src="property.value" class="w-20 h-20" />
                    </td>
                  </tr>
                </table>
              </div>
            </li>
          </ul>
        </div>

        <div v-if="credsFound.length && issuersFound.length" style="margin: 30px"></div>

        <!-- Issuers Container -->
        <div v-if="issuersFound.length">
          <p id="result-header-2">
            Found {{ issuersFound.length }} issuer{{ issuersFound.length > 1 ? 's' : '' }} who can
            issue credentials matching above criteria in your wallet:
          </p>

          <ul>
            <li
              v-for="(issuer, key) in issuersFound"
              :id="'issuer-' + key"
              :key="key"
              class="group flex flex-col items-center p-5 mb-5 w-full h-20 rounded-xl border focus-within:ring-2 focus-within:ring-offset-2 md:h-24"
            >
              <span class="text-lg font-bold">{{ issuer.name }}</span>
              <span class="text-sm md:text-base">{{ issuer.description }}</span>
            </li>
          </ul>
        </div>
      </div>

      <!-- Bottom Buttons Container -->
      <div
        class="flex sticky bottom-0 flex-row justify-between items-center py-4 px-5 w-full bg-neutrals-magnolia border-t border-neutrals-thistle"
      >
        <button id="cancelBtn" class="btn-outline" @click="cancel">
          {{ t('CHAPI.Share.decline') }}
        </button>
        <button id="share-credentials" class="btn-primary" @click="createPresentation">
          {{ t('CHAPI.Share.share') }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.notWhiteCredentialPreview:focus {
  outline: 2px solid var(--focus-color);
  outline-offset: 2px;
}
</style>
