<!--
 * Copyright SecureKey Technologies Inc. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
-->

<template>
  <div class="flex flex-col justify-start items-start w-full">
    <span v-if="heading.length" class="py-4 pl-3 text-xl font-bold text-neutrals-dark">{{
      heading
    }}</span>
    <table class="w-full border-t border-neutrals-chatelle">
      <tr
        v-for="(property, key) of credential?.properties"
        :key="key"
        class="border-b border-neutrals-thistle border-dotted"
      >
        <td class="py-4 pr-6 pl-3 w-2/5 text-neutrals-medium whitespace-nowrap">
          {{ property?.label }}
        </td>
        <td
          v-if="property?.type === 'image' || property?.schema?.contentMediaType?.includes('image')"
          class="py-4 pr-6 pl-3 w-3/5 text-neutrals-dark break-words"
        >
          <img :src="property?.value" class="w-20 h-20" />
        </td>
        <td v-else class="py-4 pr-6 pl-3 w-3/5 text-neutrals-dark break-words">
          {{ property?.value }}
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: 'CredentialDetailsTableComponent',
  props: {
    heading: {
      type: String,
      default: '',
    },
    credential: {
      type: Object,
      required: true,
    },
  },
};
</script>
