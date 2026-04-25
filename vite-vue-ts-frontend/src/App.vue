<script setup lang="ts">

  const theme = localStorage.getItem("theme");
  if (theme === "dark") {
    document.documentElement.setAttribute(
      'data-bs-theme', 'dark'
    )
  }

  const fluidLayout = localStorage.getItem("fluidLayout")

  if (fluidLayout == "true") {
    document.body.classList.add("layout-fluid");
  }
</script>

<template>
  <router-view />
</template>