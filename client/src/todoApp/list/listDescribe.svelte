<script>
    import { currentListStore, modalEditListVisible, modalDescribeListVisible } from './listStore'
    import { getListByIdService, delListByIdService } from './listService'
    import { choosedListStore } from '../todoAppStore'
    import ListEditModal from './listEditModal.svelte'
    import ListDescribeModal from './listDescribeModal.svelte'
    let modalDescribe
    modalDescribeListVisible.subscribe((value) => {
        modalDescribe = value
    })
    let modalEdit
    modalEditListVisible.subscribe((value) => {
        modalEdit = value
    })
    let choosedList
    choosedListStore.subscribe((value) => {
        choosedList = value
    })
    let list
    currentListStore.subscribe((value) => {
        list = value
    })
    function delList() {
        delListByIdService()
        .then()
        .catch()
        .finally()
    }
    $: getListByIdService(choosedList)
    .then()
    .catch()
    .finally()
</script>

<div class="list-describe">
    <table>
        <th class="title">Название</th>
        <th class="description">Описание</th>
        <th class="tools">Инструменты</th>
        <tr>
            <td class="title">
                {list.title.slice(0, 30)}
                {#if list.title.length > 30}
                    ...
                {/if}
            </td>
            <td class="description">
                {list.description.slice(0, 30)}
                {#if list.description.length > 30}
                    ...
                {/if}
            </td>
            <td class="tools">
                <button
                    id="describe"
                    on:click={() => modalDescribeListVisible.update(() => true)}
                >
                    ...
                </button>
                <button 
                    id="edit"
                    on:click={() => modalEditListVisible.update(() => true)}
                >
                    &#9998
                </button>
                <button id="del" on:click={delList}>&#128465</button>
            </td>
        </tr>
    </table>  
    {#if modalEdit}
        <ListEditModal />
    {/if}
    {#if modalDescribe}
        <ListDescribeModal />
    {/if}
</div>

<style>
    .list-describe {
        width: 100%;
        height: 100%;
    }
    table {
        width: 100%;
        text-align: center;
    }
    .title {
        width: 30%;
        white-space: nowrap;
    }
    .description {
        width: 50%;
        white-space: nowrap;
    }
    .tools {
        width: 20%;
    } 
    button {
        border: none;
        border-radius: 50%;
        min-width: 40px;
        font-size: 16px;
    }
    #describe:hover {
        background-color: rgb(241, 193, 35, 0.6);
        cursor: pointer;
    }
    #edit:hover {
        background-color: rgba(247, 244, 96, 0.6);
        cursor: pointer;
    }
    #del:hover {
        background-color: rgb(250, 7, 7, 0.6);
        cursor: pointer;
    }
</style>