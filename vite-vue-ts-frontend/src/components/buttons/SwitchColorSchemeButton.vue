<script setup lang="ts">
    import { useI18n } from "vue-i18n";
    import { NButton, NTooltip } from 'naive-ui';
    import { IconMoon, IconSun } from '@tabler/icons-vue';
    import { useColorSchemeStore } from '../../stores/colorScheme';

    const { t } = useI18n();

    interface SwitchColorSchemeButtonProps {
        size?: number,
    };

    withDefaults(defineProps<SwitchColorSchemeButtonProps>(), {
        size: 20
    });

    const colorSchemeStore = useColorSchemeStore();
</script>

<template>
    <n-tooltip trigger="hover">
        <template #trigger>
            <n-button quaternary @click.prevent="colorSchemeStore.toggle" @mousedown.prevent>
                <IconMoon :size="size" v-if="colorSchemeStore.light" />
                <IconSun :size="size" v-else />
            </n-button>
        </template>
        {{
            t(colorSchemeStore.light ?
                "Switch to dark mode" :
                "Switch to light mode")
        }}
    </n-tooltip>
</template>

<style lang="css" scoped></style>