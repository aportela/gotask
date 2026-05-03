<script setup lang="ts">
    import { reactive, onMounted, type CSSProperties, shallowRef } from 'vue';
    import { useI18n } from "vue-i18n";
    import { type SelectOption, NSelect } from 'naive-ui';
    import { api } from '../../composables/api';
    import { type SearchWorkspacesResponse } from '../../types/apiResponses';
    import { type AjaxStateInterface, defaultAjaxState } from "../../types/ajaxState";

    const { t } = useI18n();

    interface WorkspaceSelectorProps {
        style?: string | CSSProperties;
    };

    const props = defineProps<WorkspaceSelectorProps>();

    const emit = defineEmits(['error'])

    const selectedWorkspace = defineModel<string | null>('value');

    const options = shallowRef<SelectOption[]>([]);

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    const onRefresh = () => {
        Object.assign(state, defaultAjaxState);
        state.ajaxRunning = true;
        api.workspace.search().then((response: SearchWorkspacesResponse) => {
            options.value = response.data.workspaces.map((item) => ({
                value: item.id,
                label: item.name,
                disabled: false
            }))
        }).catch((error: any) => {
            emit("error", { error: error });
        }).finally(() => {
            state.ajaxRunning = false;
        });
    };

    onMounted(() => {
        onRefresh();
    });
</script>

<template>
    <n-select :placeholder="t('Workspace')" v-model:value="selectedWorkspace" :options="options" :style="style"
        :disabled="state.ajaxRunning" :loading="state.ajaxRunning"></n-select>
</template>

<style lang="css" scoped></style>