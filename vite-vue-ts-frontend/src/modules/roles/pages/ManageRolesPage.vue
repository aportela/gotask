<script setup lang="ts">
    import { useI18n } from "vue-i18n";

    import { NCard } from "naive-ui";

    const { t } = useI18n();
</script>

<template>
    <n-card :title="t('Manage roles')">
    </n-card>
</template>

<style lang="css"></style>