<!--
 * Copyright SecureKey Technologies Inc. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
-->

<template>
  <div class="flex justify-center items-start w-screen h-screen">
    <div
      class="flex flex-col justify-center items-center w-full max-w-md h-auto bg-gray-light border border-t-0 border-neutrals-black"
    >
      <div class="flex flex-col justify-start items-center py-16 px-5">
        <img src="@/assets/img/icons-error.svg" />
        <span class="mt-5 mb-3 text-xl font-bold text-center text-neutrals-dark"
          >Access Denied</span
        >
        <span class="text-lg text-center text-neutrals-medium">
          You don't have authorization to access this page, please sign in to your wallet and try
          again.</span
        >
      </div>
      <div
        class="flex flex-row justify-center items-center py-4 px-5 w-full bg-neutrals-magnolia border-t border-neutrals-thistle"
      >
        <StyledButtonComponent type="btn-primary" :loading="loading" @click="signIn()"
          >Sign in</StyledButtonComponent
        >
      </div>
    </div>
  </div>
</template>

<script>
import StyledButtonComponent from '@/components/StyledButton/StyledButtonComponent.vue';

export default {
  name: 'BlockNoAuthPage',
  components: {
    StyledButtonComponent,
  },
  data() {
    return {
      redirect: this.$route.params['redirect']
        ? this.$router.resolve(this.$route.params['redirect'])
        : '/',
      loading: false,
    };
  },
  methods: {
    signIn: function () {
      this.loading = true;
      window.open(this.redirect.href, '_blank');
    },
  },
};
</script>
