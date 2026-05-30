<script setup lang="ts">
    import { ref, computed, watch, nextTick } from 'vue';
    import { useRoute, useRouter } from 'vue-router';

    import { NTabs, NTabPane, type TabsInst } from 'naive-ui';

    import ProjectMetadataTab from '../components/ProjectMetadataTab.vue';
    import ProjectTasksTab from '../components/ProjectTasksTab.vue';
    import ProjectPermissionsTab from '../components/ProjectPermissionsTab.vue';
    import ProjectAttachmentsTab from '../components/ProjectAttachmentsTab.vue';
    import ProjectNotesTab from '../components/ProjectNotesTab.vue';
    import ProjectHistoryOperationsTab from '../components/ProjectHistoryOperationsTab.vue';

    const route = useRoute();
    const router = useRouter();

    const projectId = route.params.id as string

    // TODO: set tab with type (type tab = "metadata" | "permissions"....)
    const tab = computed({
        get: () => route.params.tab as string,
        set: (value: string) => {
            router.push({
                name: 'projectTab',
                params: {
                    id: route.params.id,
                    tab: value
                }
            });
        }
    });

    const permissionCount = ref<number>(0);
    const noteCount = ref<number>(0);
    const attachmentCount = ref<number>(0);
    const historyOperationCount = ref<number>(0);
    const taskCount = ref<number>(0);

    const permissionsTabLabel = computed(() => permissionCount.value > 0 ? `Permissions (${permissionCount.value})` : 'Permissions')
    const attachmentsTabLabel = computed(() => attachmentCount.value > 0 ? `Attachments (${attachmentCount.value})` : 'Attachments')
    const notesTabLabel = computed(() => noteCount.value > 0 ? `Notes (${noteCount.value})` : 'Notes')
    const historyTabLabel = computed(() => historyOperationCount.value > 0 ? `History (${historyOperationCount.value})` : 'History')
    const tasksTabLabel = computed(() => taskCount.value > 0 ? `Tasks (${taskCount.value})` : 'Tasks')

    const tabsRef = ref<TabsInst>();


    // recalc bar position on dynamic tab labels changes
    watch(
        () => [permissionsTabLabel.value, attachmentsTabLabel.value, notesTabLabel.value, historyTabLabel.value, tasksTabLabel.value],
        async () => {
            await nextTick();
            tabsRef.value?.syncBarPosition();
        }
    );
</script>

<template>
    <n-tabs placement="top" type="line" ref="tabsRef" animated v-model:value="tab">
        <n-tab-pane name="metadata" display-directive="show" key="metadata">
            <template #tab>
                Metadata
            </template>
            <ProjectMetadataTab mode="add" :project-id="projectId" v-model:permission-count="permissionCount"
                v-model:note-count="noteCount" v-model:attachment-count="attachmentCount"
                v-model:history-operation-count="historyOperationCount" v-model:task-count="taskCount" />
        </n-tab-pane>
        <n-tab-pane name="permissions" display-directive="show:lazy">
            <template #tab>
                {{ permissionsTabLabel }}
            </template>
            <ProjectPermissionsTab :project-id="projectId" v-model:item-count="permissionCount" />
        </n-tab-pane>
        <n-tab-pane name="notes" display-directive="show:lazy">
            <template #tab>
                {{ notesTabLabel }}
            </template>
            <ProjectNotesTab :project-id="projectId" v-model:item-count="noteCount" />
        </n-tab-pane>
        <n-tab-pane name="attachments" :tab="attachmentsTabLabel" display-directive="show:lazy"
            :key="attachmentsTabLabel">
            <ProjectAttachmentsTab :project-id="projectId" v-model:item-count="attachmentCount" />
        </n-tab-pane>
        <n-tab-pane name="history" :tab="historyTabLabel" display-directive="show:lazy" :key="historyTabLabel">
            <ProjectHistoryOperationsTab :project-id="projectId" v-model:item-count="historyOperationCount" />
        </n-tab-pane>
        <n-tab-pane name="tasks" :tab="tasksTabLabel" display-directive="show:lazy" :key="tasksTabLabel">
            <ProjectTasksTab :project-id="projectId" v-model:item-count="taskCount" />
        </n-tab-pane>
    </n-tabs>
</template>

<style lang="css" scoped>
    .avatar {
        margin-right: 4px;
    }
</style>