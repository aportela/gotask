<script setup lang="ts">
    import { computed } from "vue";
    import { useI18n } from "vue-i18n";

    import { NSelect } from 'naive-ui';
    import type { SelectMixedOption } from "naive-ui/es/select/src/interface";
    import type { SelectSize } from "naive-ui";

    import { UserTypeFilterValue, type UserTypeFilter } from "../types/filter-user-type";

    interface Props {
        size?: SelectSize,
        placeholder?: string;
        clearable?: boolean;
    }

    const { t } = useI18n();

    const props = withDefaults(defineProps<Props>(), {
        size: "medium",
        clearable: false,
    });

    const options = computed<SelectMixedOption[]>(() => [
        { label: t("All users"), value: UserTypeFilterValue.None },
        { label: t("Only administrators"), value: UserTypeFilterValue.OnlyAdministrators },
        { label: t("Only users"), value: UserTypeFilterValue.OnlyUsers }
    ]);


    const model = defineModel<UserTypeFilter>({
        default: UserTypeFilterValue.None,
    });
</script>

<template>
    <n-select :size="props.size" :options="options" v-model:value="model" :placeholder="props.placeholder"
        :clearable="props.clearable" />
</template>

<style lang="css" scoped></style>