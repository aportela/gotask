<script setup lang="ts">
    import { NCard, NInput, NFlex, NButton } from 'naive-ui';
    import { IconCancel, IconPlus, IconTrash } from '@tabler/icons-vue';
    import { computed, onMounted } from 'vue';
    //import { v7 as uuidv7 } from 'uuid';
    import type { EntityAction } from '../../types/common';

    import type { CSSProperties } from 'vue';
    import { api } from '../../composables/api';

    const emit = defineEmits(['add', 'update', 'delete', 'cancel'])

    interface ProjectTypeFormProps {
        mode: EntityAction;
        projectTypeId?: string;
        style?: string | CSSProperties;
    }

    const props = defineProps<ProjectTypeFormProps>();

    const addMode = computed<boolean>(() => props.mode === "add");
    const updateMode = computed<boolean>(() => props.mode === "update");
    const deleteMode = computed<boolean>(() => props.mode === "delete");

    const title = computed<string>(() => {
        switch (props.mode) {
            case "add":
                return "Add new project type";
            case "update":
                return "Update project type";
            case "delete":
                return "Delete project type";
            default:
                return "";
        }
    });

    const onAdd = () => {
        emit('add')
    };

    const onUpdate = () => {
        emit('update')
    };

    const onDelete = () => {
        emit('delete')
    };

    const onCancel = () => {
        emit('cancel')
    }

    /*
const onAddProjectType = () => {
projectTypes.value.push(
    new ProjectType({ id: uuidv7(), name: "" })
);
nextTick(() => {
    if (tableFooter.value) {
        tableFooter.value?.scrollIntoView({
            behavior: 'smooth',
            block: 'end'
        });
    }
});
};
*/

    const getProjectType = (id: string) => {
        api.projectTypes.get(id);
    };

    onMounted(() => {
        if (props.mode == "update" || props.mode == "delete") {
            if (props.projectTypeId) {
                getProjectType(props.projectTypeId);
            } else {
                console.error("TODO")
            }
        }
    });

</script>

<template>
    <n-card :title="title" :style="style">
        <n-input placeholder="Project type name"></n-input>
        <template #action>
            <n-flex>
                <n-button @click="onAdd" v-if="addMode">
                    <template #icon>
                        <IconPlus />
                    </template>
                    Add new
                </n-button>
                <n-button @click="onUpdate" v-else-if="updateMode">
                    <template #icon>
                        <IconPlus />
                    </template>
                    Update
                </n-button>
                <n-button @click="onDelete" v-else-if="deleteMode">
                    <template #icon>
                        <IconTrash />
                    </template>
                    Delete
                </n-button>
                <n-button @click="onCancel">
                    <template #icon>
                        <IconCancel />
                    </template>
                    Cancel
                </n-button>
            </n-flex>
        </template>
    </n-card>
</template>

<style lang="css" scoped></style>