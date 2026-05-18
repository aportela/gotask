<script setup lang="ts">
    import { computed } from "vue";
    import { useI18n } from "vue-i18n";

    import { NSelect } from 'naive-ui';
    import type { SelectMixedOption } from "naive-ui/es/select/src/interface";
    import type { SelectSize } from "naive-ui";

    import { UserAdminPermissionFilterValue, type UserAdminPermissionFilter } from "../types/user-admin-permission-filter";

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
        { label: t("userAdminPermissionSelectorValueAll"), value: UserAdminPermissionFilterValue.All },
        { label: t("userAdminPermissionSelectorValueOnlyAdministrators"), value: UserAdminPermissionFilterValue.OnlyAdministrators },
        { label: t("userAdminPermissionSelectorValueOnlyUsers"), value: UserAdminPermissionFilterValue.OnlyUsers }
    ]);


    const model = defineModel<UserAdminPermissionFilter>({
        default: UserAdminPermissionFilterValue.All,
    });
</script>

<template>
    <n-select :size="props.size" :options="options" v-model:value="model" :placeholder="props.placeholder"
        :clearable="props.clearable" />
</template>

<style lang="css" scoped></style>