<script>
    import { 
        modalDescribeTodoVisible,
        currentTodoStore
    } from './todoStore'
    import { getTodoByIdService } from './todoService'
    let currentTodo
    currentTodoStore.subscribe((value) => currentTodo = value)
    let flagStart = false
    getTodoByIdService()
    .then((value) => flagStart = value)
    .catch()
    .finally()
</script>

{#if flagStart}
    <div class="modal">
        <div class="modal-content">
            <span class="modal-name">Подробности задачи</span>
            <div class="modal-form">
                Название
                <p>{currentTodo.title}</p>
                Описание
                <p>{currentTodo.description}</p>
                Статус важности
                <p>
                    {#if currentTodo.importance}
                        Важно
                    {:else}
                        Не важно
                    {/if}
                </p>
                Статус выполнения
                <p>
                    {#if currentTodo.done}
                        Выполнено
                    {:else}
                        Не выполнено
                    {/if}
                </p>    
                <button 
                    id="cancel"
                    on:click={() => modalDescribeTodoVisible.update(() => false)}
                >
                    Вернуться
                </button>
            </div>
        </div>
    </div>
{/if}

<style>
    .modal {
        display: block;
        position: fixed;
        z-index: 1;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        overflow: auto;
        background-color: rgb(0, 0, 0, 0.8); 
    }
    .modal-content {
        background-color: #fefefe;
        margin: 5% auto;
        padding: 5px;
        border: 1px solid #888;
        width: 40%;
        min-width: 270px;
        max-width: 40%;
    }
    .modal-name {
        font-size: 20px;
        display: block;
        font-weight: bold;
    }
    .modal-form {
        margin-top: 10px;
        display: block;
        justify-content: center;
    }
    button {
        border: none;
        margin: 5px;
    }
    button:hover {
        cursor: pointer;
    }
    #cancel:hover {
        background-color: rgb(243, 23, 23, 0.8);
    }
    p {
        white-space: normal;
        overflow: auto;
        display: block;
        border: solid 1px blue;
    }
</style>