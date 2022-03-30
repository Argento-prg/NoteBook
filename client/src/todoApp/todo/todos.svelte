<script>
    import TodoCreateModal from './todoCreateModal.svelte'
    import TodoEditModal from './todoEditModal.svelte'
    import TodoDescribeModal from './todoDescribeModal.svelte'
    import { choosedListStore } from '../todoAppStore'
    import { 
        todosStore,
        choosedTodoStore,
        modalCreateTodoVisible,
        modalEditTodoVisible,
        modalDescribeTodoVisible
    } from './todoStore'
    import { 
        getAllTodosService, 
        updateDoneTodoService, 
        updateImportanceTodoService,
        deleteTodoService 
    } from './todoService'
    let choosedList
    choosedListStore.subscribe(value => choosedList = value)
    //привязка к хранилищу задач
    let todosNotDone
    let todosDone
    todosStore.subscribe(value => todosNotDone = value.filter(todo => !todo.done))
    todosStore.subscribe(value => todosDone = value.filter(todo => todo.done))
    //привязка к хранилищу показа модальных окон
    let modalCreate
    let modalEdit
    let modalDescribe
    modalCreateTodoVisible.subscribe(value => modalCreate = value)
    modalEditTodoVisible.subscribe(value => modalEdit = value)
    modalDescribeTodoVisible.subscribe(value => modalDescribe = value)
    //установка флага "Выполнено"
    function setFieldDone(value) {
        updateDoneTodoService(value)
        .then()
        .catch()
        .finally()
    }
    //установка флага "Важность"
    function setFieldImportance(value) {
        updateImportanceTodoService(value)
        .then()
        .catch()
        .finally()
    }
    //удалить задачу
    function delTodo() {
        deleteTodoService()
        .then()
        .catch()
        .finally()
    }
    let flagStart = false
    //получаем все задачи при загрузке компонента
    $: getAllTodosService()
    .then((value) => {
        choosedList = choosedList//для реактивности элемента
        flagStart = value
    })
    .catch()
    .finally()
</script>

{#if flagStart && choosedList !== -1}
    <div class="todos-main">
        <div class="todos-header">
            <div class="todos-header-not-done">
                Ожидают выполнения
            </div>
            <div class="todos-header-done">
                Выполнено
            </div>
        </div>
        <div class="todos-content">
            <div class="todos-list-not-done">
                {#each todosNotDone as todo (todo.id)}
                    <div class="todo">
                        <div class="left-tools">
                            <button
                                id="flag-done"
                                on:click={() => {
                                    choosedTodoStore.update(() => todo.id)
                                    setFieldDone(true)
                                }}
                            >
                                &#9898
                            </button>
                            {#if todo.importance}
                                <button
                                    id="flag-importance"
                                    on:click={() => {
                                        choosedTodoStore.update(() => todo.id)
                                        setFieldImportance(false)
                                    }}
                                >
                                    &#128276
                                </button>
                            {:else}
                                <button
                                    id="flag-importance"
                                    on:click={() => {
                                        choosedTodoStore.update(() => todo.id)
                                        setFieldImportance(true)
                                    }}
                                >
                                    &#128277
                                </button>
                            {/if}
                        </div>
                        <div class="central">
                            {todo.title.slice(0, 25)}
                            {#if todo.title.length > 25}
                                ...
                            {/if}
                        </div>
                        <div class="right-tools">
                            <button
                                id="describe"
                                on:click={() => {
                                    choosedTodoStore.update(() => todo.id)
                                    modalDescribeTodoVisible.update(() => true)
                                }}
                            >
                                ...
                            </button>
                            <button
                                id="edit"
                                on:click={() => {
                                    choosedTodoStore.update(() => todo.id)
                                    modalEditTodoVisible.update(() => true)
                                }}
                            >
                                &#9998
                            </button>
                            <button
                                id="del"
                                on:click={() => {
                                    choosedTodoStore.update(() => todo.id)
                                    delTodo()
                                }}
                            >
                                &#128465
                            </button>
                        </div>
                    </div>
                {/each}
            </div>
            <div class="todos-list-done">
                {#each todosDone as todo (todo.id)}
                    <div class="todo">
                        <div class="left-tools">
                            <button
                                id="flag-done"
                                on:click={() => {
                                    choosedTodoStore.update(() => todo.id)
                                    setFieldDone(false)
                                }}
                            >
                                &#9989
                            </button>
                            {#if todo.importance}
                                <button
                                    id="flag-importance"
                                    on:click={() => {
                                        choosedTodoStore.update(() => todo.id)
                                        setFieldImportance(false)
                                    }}
                                >
                                    &#128276
                                </button>
                            {:else}
                                <button
                                    id="flag-importance"
                                    on:click={() => {
                                        choosedTodoStore.update(() => todo.id)
                                        setFieldImportance(true)
                                    }}
                                >
                                    &#128277
                                </button>
                            {/if}
                        </div>
                        <div class="central">
                            {todo.title.slice(0, 25)}
                            {#if todo.title.length > 25}
                                ...
                            {/if}
                        </div>
                        <div class="right-tools">
                            <button
                                id="describe"
                                on:click={() => {
                                    choosedTodoStore.update(() => todo.id)
                                    modalDescribeTodoVisible.update(() => true)
                                }}
                            >
                                ...
                            </button>
                            <button
                                id="del"
                                on:click={() => {
                                    choosedTodoStore.update(() => todo.id)
                                    delTodo()
                                }}
                            >
                                &#128465
                            </button>
                        </div>
                    </div>
                {/each}
            </div>
        </div>
        <hr />
        <div class="button-create">
            <button
                id="create-btn"
                on:click={() => modalCreateTodoVisible.update(() => true)}
            >
                Создать
            </button>
        </div>
        {#if modalCreate}
            <TodoCreateModal />
        {/if}
        {#if modalEdit}
            <TodoEditModal />
        {/if}
        {#if modalDescribe}
            <TodoDescribeModal />
        {/if}
    </div>
{/if}

<style>
    .todos-main {
        height: 100%;
        width: 100%;
        background-color: aliceblue;
        text-align: center;
        margin-right: auto;
        margin-left: auto;
    }
    .todos-header {
        display: flex;
        font-weight: bold;
        height: 5%;
    }
    .todos-header-not-done {
        width: 50%;
        border-right: gray 1px inset;
    }
    .todos-header-done {
        width: 50%;
    }
    .todos-content {
        display: flex;
        max-height: 85%;
    }
    .todos-list-not-done {
        width: 50%;
        overflow-x: auto;
        overflow-y: auto;
        max-height: 100%;
        display: block;
        border-right: gray 1px inset;
    }
    .todos-list-done {
        width: 50%;
        overflow-x: auto;
        overflow-y: auto;
        max-height: 100%;
        display: block;
    }
    hr {
        margin: 0;
        padding: 0;
        height: 1%;
    }
    button {
        border: none;
        border-radius: 25px;
        font-size: 12px;
        padding: 4px;
        margin-top: 5px;
        margin-bottom: 5px;
    }
    .button-create {
        height: 9%;
    }
    #create-btn{
        background-color: rgb(183, 247, 35, 0.8);
        font-weight: bold;
        cursor: pointer;
    }
    #create-btn:hover {
        background-color: rgb(143, 243, 42, 0.8);
        font-weight: bold;
        cursor: pointer;
    }
    .todo {
        display: flex;
        margin-top: 10px;
    }
    .left-tools {
        width: 20%;
    }
    .central {
        width: 60%;
        white-space: nowrap;
    }
    .right-tools {
        width: 20%;
    }
    #flag-done:hover {
        background-color: rgba(36, 204, 247, 0.6);
        cursor: pointer;
    }
    #flag-importance:hover {
        background-color: rgba(243, 247, 20, 0.6);
        cursor: pointer;
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