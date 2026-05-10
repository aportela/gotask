<script setup lang="ts">
    import { onMounted, onBeforeUnmount, ref, reactive, computed, shallowRef, h, type Component } from 'vue';
    import { useI18n } from "vue-i18n";
    import { NAvatar, NInput, NSelect, NIcon, NButton, NModal, NButtonGroup, useDialog } from 'naive-ui';
    import { IconUser, IconUserKey, IconSearch, IconPlus, IconEdit, IconTrash, IconTrashOff } from '@tabler/icons-vue';
    import { api } from '../composables/api';
    import { type AjaxStateInterface, defaultAjaxState } from '../types/ajaxState';
    import type { AxiosAPIError } from '../composables/axios';
    import { type EntityAction } from '../types/common';
    import { useLoadingStore } from '../stores/loading';
    import { useSessionStore } from '../stores/session';
    import { useNotify } from '../composables/notification';
    import { default as UserForm } from '../components/forms/UserForm.vue';
    import { default as ManageTable } from '../components/custom/ManageTable.vue';
    import { default as DateFilter } from '../components/forms/DateFilter.vue';
    import { useAppBus, type AppBusEvent } from '../composables/bus';
    import { userService } from '../api/services/user';
    import { User } from '../api/models/user';
    import type { UserResponse } from '../api/types/dto/user';

    const appBus = useAppBus();

    const { notify } = useNotify();

    const { t } = useI18n();

    const sessionStore = useSessionStore();

    const loadingStore = useLoadingStore();

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    class SearchableUser extends User {
        _searchName: string;
        _searchEmail: string;

        constructor(data: UserResponse) {
            super(data);
            this._searchName = data.name.toLowerCase();
            this._searchEmail = data.email.toLowerCase();
        }
    }

    const users = shallowRef<SearchableUser[]>([]);

    const filterUserOptions = [
        { label: 'All users', value: 0 },
        { label: 'Administrators only', value: 1 },
        { label: 'Users only', value: 2 }
    ];

    const filterByUsername = ref<string | null>(null);
    const filterByEmail = ref<string | null>(null);
    const userFilterType = ref<number | null>(0);

    const filteredUsers = computed<SearchableUser[]>(() => {
        const name = filterByUsername.value?.trim().toLowerCase();
        const email = filterByEmail.value?.trim().toLowerCase();
        return users.value.filter(u => {
            const matchName = !name || u._searchName.includes(name);
            const matchEmail = !email || u._searchEmail.includes(email);
            const matchType = userFilterType.value === 0 || (u.isSuperUser && userFilterType.value == 1) || (!u.isSuperUser && userFilterType.value == 2);
            return matchName && matchEmail && matchType;
        });
    });

    const onRefresh = async () => {
        state.ajaxRunning = true;
        loadingStore.set(true);
        try {
            const response = await userService.search();
            users.value = response.users.map((user: UserResponse) => new SearchableUser(user));
        }
        finally {
            state.ajaxRunning = false;
            loadingStore.set(false);
        }
        /*
        api.user.search().then((successResponse: SearchUsersResponse) => {
            users.value = successResponse.data.users.map((u: UserInterface) => new UserClass(u));
        }).catch((errorResponse: AxiosAPIError) => {
            state.ajaxErrors = true;
            if (errorResponse.isAPIError) {
                state.ajaxAPIErrorDetails = errorResponse.customAPIErrorDetails;
                switch (errorResponse.status) {
                    case 401:
                        state.ajaxErrors = false;
                        appBus.emitReauthRequired("ManageUsersPage.onRefresh");
                        break;
                    default:
                        state.ajaxErrorMessage = t("There was a problem while refreshing the user list");
                        break;
                }
            } else {
                state.ajaxErrorMessage = t("There was a problem while refreshing the user list");
                console.error(errorResponse);
            }
        }).finally(() => {
            loadingStore.set(false);
            state.ajaxRunning = false;
        });
        */
    };

    const selectedUserId = ref<string | undefined>(undefined);

    const onAddUser = () => {
        actionDialogMode.value = "add";
    };

    const onUpdateUser = (user: User, _index: number) => {
        actionDialogMode.value = "update";
        selectedUserId.value = user.id;
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

    const onAdd = () => {
        isVisibleActionDialog.value = false;
        notify('success', t("User added"))
        onRefresh();
    };

    const onUpdate = () => {
        isVisibleActionDialog.value = false;
        notify('success', t("User updated"))
        onRefresh();
    };

    const onCancel = () => {
        isVisibleActionDialog.value = false;
    };

    onMounted(() => {
        onRefresh();
        appBus.on((event: AppBusEvent) => {
            if (event.type == "reauthValidNotify" && event.to.includes("ManageUsersPage.onRefresh")) {
                onRefresh();
            }
        });
    });

    onBeforeUnmount(() => {
        // TODO: manage this
        //bus.off("reAuthSucess");
    });


    const onDelete = (userId: string) => {
        Object.assign(state, defaultAjaxState);
        state.ajaxRunning = true;
        api.user.delete(userId).then((_response: any) => {
            notify('success', t("User deleted"))
            onRefresh();
        }).catch((errorResponse: AxiosAPIError) => {
            state.ajaxErrors = true;
            if (errorResponse.isAPIError) {
                state.ajaxAPIErrorDetails = errorResponse.customAPIErrorDetails;
                switch (errorResponse.status) {
                    case 401:
                        state.ajaxErrors = false;
                        appBus.emitReauthRequired("ManageUsersPage.onDelete");
                        break;
                    case 404:
                        state.ajaxErrorMessage = t("We couldn’t find the specified user");
                        break;
                    default:
                        state.ajaxErrorMessage = t("There was a problem deleting the user data");
                        break;
                }
            } else {
                state.ajaxErrorMessage = t("There was a problem deleting the user data");
                console.error(errorResponse);
            }
        }).finally(() => {
            state.ajaxRunning = false;
        });
    };

    const onUnDelete = (userId: string) => {
        Object.assign(state, defaultAjaxState);
        state.ajaxRunning = true;
        api.user.unDelete(userId).then((_response: any) => {
            notify('success', t("User undeleted"))
            onRefresh();
        }).catch((errorResponse: AxiosAPIError) => {
            state.ajaxErrors = true;
            if (errorResponse.isAPIError) {
                state.ajaxAPIErrorDetails = errorResponse.customAPIErrorDetails;
                switch (errorResponse.status) {
                    case 401:
                        state.ajaxErrors = false;
                        appBus.emitReauthRequired("ManageUsersPage.onUnDelete");
                        break;
                    case 404:
                        state.ajaxErrorMessage = t("We couldn’t find the specified user");
                        break;
                    default:
                        state.ajaxErrorMessage = t("There was a problem undeleting the user data");
                        break;
                }
            } else {
                state.ajaxErrorMessage = t("There was a problem undeleting the user data");
                console.error(errorResponse);
            }
        }).finally(() => {
            state.ajaxRunning = false;
        });
    };

    // TODO: this is already declared on menu.ts

    const renderIcon = (icon: Component) => {
        return (size = 32) =>
            () =>
                h(
                    NIcon,
                    { size },
                    {
                        default: () => h(icon),
                    },
                );
    };


    const dialog = useDialog()

    function confirmDelete(user: User, _index: number) {
        dialog.warning({
            title: t("Delete user"),
            icon: renderIcon(IconTrash)(24),
            content: () =>
                h('div', [
                    `You are about to delete the user "${user.name}" from the system.`,
                    h('br'),
                    h('br'),
                    'Do you want to continue?',
                ]),
            positiveText: t("Delete"),
            negativeText: t("Cancel"),
            onPositiveClick: () => {
                onDelete(user.id);
            },
        })
    }

    function confirmUnDelete(user: User, _index: number) {
        dialog.warning({
            title: t("Restore user"),
            icon: renderIcon(IconTrashOff)(24),
            content: () =>
                h('div', [
                    `You are about to restore the user "${user.name}" from the system.`,
                    h('br'),
                    h('br'),
                    'Do you want to continue?',
                ]),
            positiveText: t("Restore"),
            negativeText: t("Cancel"),
            onPositiveClick: () => {
                onUnDelete(user.id);
            },
        })
    }
</script>

<template>
    <n-modal v-model:show="isVisibleActionDialog">
        <UserForm :mode="actionDialogMode" :user-id="selectedUserId" style="width: 40%;" @add="onAdd" @update="onUpdate"
            @delete="onDelete" @cancel="onCancel" @undelete="onUnDelete" />
    </n-modal>

    <ManageTable size="small" title="Manage users">
        <template #thead>
            <tr>
                <th class="text-center">Type</th>
                <th>Name</th>
                <th>Email</th>
                <th class="hide-mobile">Created at</th>
                <th class="hide-mobile">Updated at</th>
                <th class="hide-mobile">Deleted at</th>
                <th class="text-center">Operations</th>
            </tr>
            <tr class="hide-mobile">
                <th>
                    <n-select size="small" trigger="click" :options="filterUserOptions" v-model:value="userFilterType"
                        placeholder="Search by user type">
                    </n-select>
                </th>
                <th>
                    <n-input size="small" placeholder="search by name..." v-model:value="filterByUsername" clearable>
                        <template #prefix>
                            <n-icon>
                                <IconSearch />
                            </n-icon>
                        </template>
                    </n-input>
                </th>
                <th>
                    <n-input size="small" placeholder="search by email..." v-model:value="filterByEmail" clearable>
                        <template #prefix>
                            <n-icon>
                                <IconSearch />
                            </n-icon>
                        </template>
                    </n-input>
                </th>
                <th>
                    <DateFilter />
                </th>
                <th>
                    <DateFilter />
                </th>
                <th>
                    <DateFilter />
                </th>
                <th class="text-center">
                    <n-button size="small" block @click="onAddUser">
                        <template #icon>
                            <NIcon>
                                <IconPlus />
                            </NIcon>
                        </template>
                        {{ t("Add") }}
                    </n-button>
                </th>
            </tr>
        </template>
        <template #tbody>
            <tr v-for="user, index in filteredUsers" :key="user.id">
                <td class="text-center">
                    <IconUserKey v-if="user.isSuperUser" color="red" />
                    <IconUser v-else />
                </td>
                <td>
                    <div style="display: flex; align-items: center; gap: 8px;">
                        <n-avatar v-if="user.avatarURL" :src="user.avatarURL" class="avatar" /> {{ user.name }}
                    </div>
                </td>
                <td><a :href="'mailto:' + user.email">{{ user.email }}</a></td>
                <td class="hide-mobile">{{ user.createdAt.toLocaleString() }}</td>
                <td class="hide-mobile">{{ user.updatedAt?.toLocaleString() }}</td>
                <td class="hide-mobile">{{ user.deletedAt?.toLocaleString() }}</td>
                <td class="text-center">
                    <n-button-group v-if="!user.deletedAt">
                        <n-button size="small" @click="onUpdateUser(user, index)">
                            {{ t("Update") }}
                            <template #icon>
                                <IconEdit :size="22" />
                            </template>
                        </n-button>
                        <n-button size="small" @click="confirmDelete(user, index)"
                            :disabled="user.id === sessionStore.sessionUserId">
                            {{ t("Delete") }}
                            <template #icon>
                                <IconTrash :size="22" />
                            </template>
                        </n-button>

                    </n-button-group>
                    <n-button size="small" @click="confirmUnDelete(user, index)" v-else>
                        {{ t("Restore") }}
                        <template #icon>
                            <IconTrashOff :size="22" />
                        </template>
                    </n-button>
                </td>
            </tr>
        </template>
    </ManageTable>
</template>

<style lang="css">
    .avatar {
        margin-right: 4px;
    }

    @media (max-width: 768px) {
        .hide-mobile {
            display: none;
        }
    }
</style>