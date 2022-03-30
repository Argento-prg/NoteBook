<script>
    import List from './list/list.svelte'
    import ListDescribe from './list/listDescribe.svelte'
    import Todos from './todo/todos.svelte'
    import { choosedListStore } from './todoAppStore'
    let choosedList
    choosedListStore.subscribe((value) => {
        choosedList = value
    })
</script>

<div class="todoApp">
    <div class="lists">
        <span class="content-head">Мои списки задач</span>
        <List />
    </div>
    <div class="todos">
        <span class="content-head">Подробности о выбранном списке</span>
        {#if choosedList !== -1}
            <div class="list-describe">
                <ListDescribe />
            </div>
            <span class="content-head">Задачи</span>
            <div class="todos-workspace">
                <Todos />
            </div>
        {/if}
    </div>
</div>

<style>
    .todoApp {
        display: flex;
        height: 95%;
        overflow: hidden;
        position: relative;
    }
    .lists {
        width: 20%;
        border-right: solid black 2px;
    }
    .todos {
        width: 80%;
        height: 100%;
    }
    .content-head {
        height: 4%;
        display: block;
        text-align: center;
        background-color: beige;
    }
    .list-describe {
        height: 15%;
    }
    .todos-workspace {
        height: 77%;
    }
</style>