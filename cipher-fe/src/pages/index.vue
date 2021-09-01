<script setup lang="ts">
const plain = ref('')
const mode = ref('')
</script>

<template>
  <div>
    <div class="flex justify-center items-center space-x-2 text-4xl mb-4">
      <heroicons-solid:key class="inline-block text-3xl" />
      <div>cipher</div>
    </div>

    <div>
      <select>
        <option value="affine">
          Affine
        </option>
        <option value="playfair">
          Playfair
        </option>
        <option value="vigenere-standard">
          Vigenere Standard
        </option>
        <option value="Affine">
          sa
        </option>
      </select>
    </div>

    <div>
      <label class="text-sm opacity-75 italic" for="plaintext">Plaintext</label>
    </div>
    <div>
      <textarea
        id="plaintext"
        v-model="plain"
        placeholder="Enter here..."
        type="text"
        autocomplete="false"
        p="x-2 y-2"
        w="250px"
        bg="transparent"
        text="center sm"
        border="~ rounded gray-200 dark:gray-700"
        outline="none active:none"
      />
    </div>
    <div class="flex flex-col">
    </div>

    <div>
      <button
        class="w-32 m-3 text-sm btn"
      >
        Go
      </button>
    </div>
  </div>
</template>
