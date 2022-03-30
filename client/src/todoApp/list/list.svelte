<script>
    import { getAllListsService } from './listService'
    import { choosedListStore } from '../todoAppStore'
    import { modalCreateListVisible } from './listStore'
    import { listsStore } from './listStore'
    import ListCreateModal from './listCreateModal.svelte'
    let modal
    modalCreateListVisible.subscribe((value) => {
        modal = value
    })
    let choosedList
    choosedListStore.subscribe((value) => {
        choosedList = value
    })
    let lists
    listsStore.subscribe((value) => {
        lists = value
    })
    getAllListsService()
    .then()
    .catch()
    .finally()
</script>

<div class="listMain">
    <div class="allLists">
        {#each lists as list (list.id)}
            {#if list.id !== choosedList}
                <div 
                    class='list'
                    on:click={
                        () => {
                            choosedListStore.update(() => list.id)
                        }
                    }
                >
                    {list.title.slice(0, 25)}
                    {#if list.title.length > 25}
                        ...
                    {/if}
                </div>
            {:else}
                <div 
                    class='list choosed-list'
                    on:click={
                        () => {
                            choosedListStore.update(() => list.id)
                        }
                    }
                >
                    {list.title.slice(0, 25)}
                    {#if list.title.length > 25}
                        ...
                    {/if}
                </div>
            {/if}
        {/each}
    </div>
    <hr />
    <button on:click={() => modalCreateListVisible.update(() => true)}>создать</button>
    {#if modal}
        <ListCreateModal />
    {/if}
</div>

<style>
    .allLists {
        overflow-x: auto;
        overflow-y: auto;
        max-height: 85%;
    }
    .listMain {
        text-align: center;
        margin-right: auto;
        margin-left: auto;
        margin: 0 5px;
        height: 100%;
    }
    button {
        border: none;
        margin-top: 5px;
        margin-bottom: 5px;
        border-radius: 25px;
        font-size: 14px;
        padding: 4px;
        white-space: nowrap;
    }
    hr {
        margin: 0;
    }
    button:hover {
        background-color: rgb(143, 243, 42);
        font-weight: bold;
        cursor: pointer;
    }
    .list {
        border-style: inset;
        border-width: 2px;
        border-color: blue;
        border-radius: 15px;
        cursor: pointer;
        margin-top: 5px;
        font-size: 12px;        
    }
    .choosed-list {
        background-color: rgb(62, 247, 117, 0.5);
        font-weight: bold;
    }
</style>