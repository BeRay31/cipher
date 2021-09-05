<template>
  <div class="flex flex-col justify-center items-center p-12">
    <!-- INPUT-CONTAINER -->
    <div class="mt-4 w-[100%] flex flex-col justify-center items-center">
      <!-- CONTENT-CONTAINER -->
      <div class="flex flex-row justify-evenly items-center w-full mb-8">
        <!-- MENU-CONTAINER -->
        <!-- MODE-CONTAINER -->
        <cipher-selection />
        <!-- INPUT-CONTENT-CONTAINER -->
        <mode-selection />
      </div>
      <div class="w-full">
        <result />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
</script>
