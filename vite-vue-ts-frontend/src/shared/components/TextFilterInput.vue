<script setup lang="ts">
    import { NInput, NIcon } from 'naive-ui';
    import type { InputSize } from "naive-ui";
    import { IconSearch } from '@tabler/icons-vue';

    interface Props {
        size?: InputSize,
        placeholder?: string;
        clearable?: boolean;
    }

    const props = withDefaults(defineProps<Props>(), {
        size: "medium",
        clearable: false,
    });

    const model = defineModel<string>({
        default: "",
    });
</script>

<template>
    <n-input :size="props.size" v-model:value="model" :placeholder="props.placeholder" :clearable="props.clearable">
        <template #prefix>
            <n-icon>
                <IconSearch />
            </n-icon>
        </template>
    </n-input>
</template>

<style lang="css" scoped></style>