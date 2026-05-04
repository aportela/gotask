<script setup lang="ts">
    import { h, reactive, watch, onMounted, type CSSProperties, shallowRef, type VNodeChild } from 'vue';
    import { useI18n } from "vue-i18n";
    import { NSelect, NIcon } from 'naive-ui';
    import { api } from '../../composables/api';
    import { type SearchWorkspacesResponse } from '../../types/apiResponses';
    import { type AjaxStateInterface, defaultAjaxState } from "../../types/ajaxState";
    import { type NaiveUISelectOptionWithColor } from '../../types/customNaiveUI';
    import { IconSquareFilled } from '@tabler/icons-vue';
    import { useCurrentWorkspaceStore } from '../../stores/currentWorkspace';

    const { t } = useI18n();

    const currentWorkspaceStore = useCurrentWorkspaceStore();

    interface WorkspaceSelectorProps {
        style?: string | CSSProperties;
    };

    const props = defineProps<WorkspaceSelectorProps>();

    const emit = defineEmits(['error'])

    const selectedWorkspace = defineModel<string | null>('value');

    const options = shallowRef<NaiveUISelectOptionWithColor[]>([]);

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    watch(selectedWorkspace, (val) => {
        if (val) {
            currentWorkspaceStore.set(val);
        }
    });

    const findValidWorkspace = (options: NaiveUISelectOptionWithColor[], stored?: string | null) => {
        if (!options.length) return null;

        if (stored && options.some(o => o.value === stored)) {
            return stored;
        }

        return options[0].value as string;
    };

    const onRefresh = () => {
        Object.assign(state, defaultAjaxState);
        state.ajaxRunning = true;
        api.workspace.search().then((response: SearchWorkspacesResponse) => {
            options.value = response.data.workspaces.map((item) => ({
                value: item.id,
                label: item.name,
                disabled: false,
                color: item.hexColor,
            }));
            selectedWorkspace.value = findValidWorkspace(
                options.value,
                currentWorkspaceStore.currentWorkspaceId
            );
        }).catch((error: any) => {
            emit("error", { error: error });
        }).finally(() => {
            state.ajaxRunning = false;
        });
    };

    const renderLabel = (option: NaiveUISelectOptionWithColor): VNodeChild => {
        if (option.type === 'group')
            return `${option.label}(Cool!)`
        return [
            h(
                NIcon,
                {
                    color: option.color,
                    style: {
                        verticalAlign: '-0.15em',
                        marginRight: '4px'
                    }
                },
                {
                    default: () => h(IconSquareFilled)
                }
            ),
            option.label as string
        ]
    }
    onMounted(() => {
        onRefresh();
    });
</script>

<template>
    <n-select :placeholder="t('Workspace')" v-model:value="selectedWorkspace" :options="options"
        :render-label="renderLabel" :style="style" :disabled="state.ajaxRunning"
        :loading="state.ajaxRunning"></n-select>
</template>

<style lang="css" scoped></style>