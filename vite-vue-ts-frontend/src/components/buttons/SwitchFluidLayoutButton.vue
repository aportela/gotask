<script setup lang="ts">
    import { NButton } from 'naive-ui';
    import { IconLayout } from '@tabler/icons-vue';
    import { useUserSettingsStore } from '../../stores/userSettings';

    const userSettingsStore = useUserSettingsStore();

</script>

<template>
    <n-button quaternary @click="userSettingsStore.toggleFluidLayout" @mousedown.prevent>
        <IconLayout />
    </n-button>
</template>

<style lang="css" scoped></style>