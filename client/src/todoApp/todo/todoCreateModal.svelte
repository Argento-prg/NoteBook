<script>
    import { 
        modalCreateTodoVisible,
        maxLengthTitle,
        maxLengthDescription, 
    } from './todoStore'
    import { createTodoService } from './todoService'
    let title = ''
    let description = ''
    let importance = false
    function saveTodo () {
        createTodoService(title, description, importance)
        .then((value) => {
            if (value) {
                modalCreateTodoVisible.update(() => false)
            }
        })
        .catch()
        .finally()
    }
</script>

<div class="modal">
    <div class="modal-content">
        <span class="modal-name">Создание задачи</span>
        <div class="modal-form">
            <textarea 
                cols=50
                rows=2
                placeholder="Название задачи" 
                maxlength={maxLengthTitle}
                bind:value={title}
            />
            <textarea 
                cols=50
                rows=6
                placeholder="Описание задачи" 
                maxlength={maxLengthDescription}
                bind:value={description}
            />
            <div>
                <span>Отметить как важное?</span> 
                <input
                    id="importance" 
                    type="checkbox"
                    bind:checked={importance}
                />
            </div>
            <button
                id="save"
                on:click={saveTodo}
            >
                Сохранить
            </button>
            <button 
                id="cancel"
                on:click={() => modalCreateTodoVisible.update(() => false)}
            >
                Отмена
            </button>
        </div>
    </div>
</div>

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
        margin: 10% auto;
        padding: 5px;
        border: 1px solid #888;
        width: 40%;
        max-width: 40%;
        min-width: 270px;
    }
    .modal-name {
        font-size: 20px;
        display: block;
        font-weight: bold;
    }
    .modal-form {
        display: block;
        justify-content: center;
    }
    input {
        display: block;
        margin: 10px auto;
    }
    #importance {
        display: inline-block;
    }
    button {
        border: none;
        margin: 5px;
    }
    button:hover {
        cursor: pointer;
    }
    #save:hover {
        background-color: rgb(132, 245, 19, 0.8);
    }
    #cancel:hover {
        background-color: rgb(243, 23, 23, 0.8);
    }
    textarea {
        margin: 10px auto;
        display: block;
        resize: none;
    }
</style>