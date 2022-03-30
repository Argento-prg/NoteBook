<script>
    import { 
        modalDescribeListVisible,
        currentListStore
    } from './listStore'
    let currentList
    currentListStore.subscribe((value) => currentList = value)
</script>

<div class="modal">
    <div class="modal-content">
        <span class="modal-name">Подробное описание списка</span>
        <div class="modal-form">
            Название
            <p>{currentList.title}</p>
            Описание
            <p>{currentList.description}</p>    
            <button 
                id="cancel"
                on:click={() => modalDescribeListVisible.update(() => false)}
            >
                Назад
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
        margin: 5% auto;
        padding: 5px;
        border: 1px solid #888;
        width: 40%;
        min-width: 270px;
        max-width: 40%;
        text-align: center;
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
        border-radius: 15px;
    }
    button:hover {
        cursor: pointer;
    }
    #cancel:hover {
        background-color: rgb(243, 23, 23, 0.5);
    }
    p {
        white-space: normal;
        overflow: auto;
        display: block;
        border: solid 1px blue;
    }
</style>