<script>
    import { modalEditListVisible, currentListStore } from './listStore'
    import { editListService } from './listService'
    import { maxLengthTitle, maxLengthDescription } from './listStore'
    let oldListData
    currentListStore.subscribe((value) => oldListData = value)
    let title = oldListData.title
    let description = oldListData.description
    function saveList () {
        editListService(title, description)
        .then((value) => {
            if (value) {
                modalEditListVisible.update(() => false)
            }
        })
        .catch()
        .finally()
    }
</script>

<div class="modal">
    <div class="modal-content">
        <span class="modal-name">Обновление списка</span>
        <div class="modal-form">
            <textarea 
                cols=50
                rows=2
                placeholder="Название списка" 
                maxlength={maxLengthTitle}
                bind:value={title}
            />
            <textarea 
                cols=50
                rows=4
                placeholder="Описание списка" 
                maxlength={maxLengthDescription}
                bind:value={description}
            />
            <button
                id="save"
                on:click={saveList}
            >
                Сохранить
            </button>
            <button 
                id="cancel"
                on:click={() => modalEditListVisible.update(() => false)}
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
        text-align: center;
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