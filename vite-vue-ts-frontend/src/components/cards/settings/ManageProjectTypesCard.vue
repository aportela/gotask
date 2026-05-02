<script setup lang="ts">
    import { ref, reactive, onMounted, nextTick, computed } from 'vue'
    import { NSpin, NTable, NButton, NGrid, NGridItem, NFlex, useDialog, NModal, NTag } from 'naive-ui'
    import { api } from '../../../composables/api';
    import { IconDeviceFloppy, IconEdit, IconPlus, IconTrash } from '@tabler/icons-vue';
    import ProjectTypeForm from '../../forms/ProjectTypeForm.vue';
    import { getNaiveUITagColorProperty } from '../../../composables/color';
    import type { ProjectTypeInterface } from '../../../types/models/projectType';
    import { type AjaxStateInterface, defaultAjaxState } from '../../../types/ajaxState';
    import { type EntityAction } from '../../../types/common';

    const dialog = useDialog();

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    if (dialog !== null) {
        dialog.destroyAll();
    }

    const tableFooter = ref<HTMLElement | null>(null);

    console.log(tableFooter);

    nextTick(() => { });

    const projectTypes = ref<ProjectTypeInterface[]>([]);

    const onRefresh = () => {
        state.ajaxRunning = true;
        api.projectTypes.search().then((successResponse: any) => {
            projectTypes.value = successResponse.data.projectTypes;
        }).catch((errorResponse: any) => {
            console.log(errorResponse);
        }).finally(() => { state.ajaxRunning = false; })
    };



    onMounted(() => {
        onRefresh();
    });

    const selectedProjectTypeId = ref<string | undefined>(undefined);

    const onAddProjectType = () => {
        actionDialogMode.value = "add";
    };
    const onUpdateProjectType = (_projectType: ProjectTypeInterface, _index: number) => {
        actionDialogMode.value = "update";
        selectedProjectTypeId.value = _projectType.id;
    };
    const onDeleteProjectType = (_projectType: ProjectTypeInterface, _index: number) => {
        actionDialogMode.value = "delete";
        selectedProjectTypeId.value = _projectType.id;
    };

    const actionDialogMode = ref<EntityAction>("none");

    const isVisibleActionDialog = computed<boolean>({
        get: () => actionDialogMode.value !== "none",
        set: (value: boolean) => {
            if (!value) {
                actionDialogMode.value = "none";
            }
        }
    });

</script>

<template>
    <n-spin :show="state.ajaxRunning">
        <n-modal v-model:show="isVisibleActionDialog" @cancel="isVisibleActionDialog = false">
            <ProjectTypeForm :mode="actionDialogMode" :project-type-id="selectedProjectTypeId" style="width: 40%;" />
        </n-modal>
        <n-table size="small">
            <caption class="table-caption">
                <n-grid :cols="2" align="center">
                    <n-grid-item style="text-align: left;">
                        <span class="table-caption-title">Project types</span>
                    </n-grid-item>
                    <n-grid-item style="display: flex; justify-content: flex-end;">
                        <n-flex>
                            <n-button @click="onAddProjectType">
                                <template #icon>
                                    <IconPlus />
                                </template>
                                Add new
                            </n-button>
                            <n-button disabled>
                                <template #icon>
                                    <IconDeviceFloppy />
                                </template>
                                Save
                            </n-button>
                        </n-flex>
                    </n-grid-item>
                </n-grid>
            </caption>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="projectType, index in projectTypes" :key="projectType.id">
                    <td><n-tag :color="getNaiveUITagColorProperty(projectType.hexColor)">{{ projectType.name
                    }}</n-tag></td>
                    <td class="text-center">
                        <n-flex>
                            <n-button @click="onUpdateProjectType(projectType, index)">
                                Update
                                <template #icon>
                                    <IconEdit :size="22" />
                                </template>
                            </n-button>
                            <n-button @click="onDeleteProjectType(projectType, index)">
                                Delete
                                <template #icon>
                                    <IconTrash :size="22" />
                                </template>
                            </n-button>
                        </n-flex>
                    </td>
                </tr>
            </tbody>
            <tfoot ref="tableFooter"></tfoot>
        </n-table>
    </n-spin>
</template>

<style lang="css" scoped>
    .table-caption {
        padding-bottom: 4px;
    }

    .table-caption-title {
        font-weight: 700;
        font-size: large;

    }

    .text-center {
        text-align: center;
    }
</style>