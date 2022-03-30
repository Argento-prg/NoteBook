<script>
    import TodoApp from './todoApp/todoApp.svelte'
    import HomeMain from './homeMain.svelte'
    import { isLoggedStore, store } from './storeApp'

    let choiceApp = 0
    function exit() {
        store.deleteName()
        store.deleteToken()
        isLoggedStore.update(value => false)
    }
</script>

<div class="workspace">
    <nav>
        <button id="logo" on:click={() => choiceApp = 0}>NoteBook</button>
        <div class="appList">
            <button on:click={() => choiceApp = 1}>Записная книжка</button>
        </div>
        <button id="exit" on:click={exit}>Выход</button>
    </nav>
    {#if choiceApp === 0}
        <HomeMain />
    {:else if choiceApp === 1}
        <TodoApp />
    {/if}
</div>

<style>
    .workspace {
        width: 100%;
        height: 90%;
        background: linear-gradient(30deg, #69b7eb, #b3dbd3);
    }
    nav {
        display: flex;
        height: 5%;
        margin: 0;
        padding: 0;
        background-color: paleturquoise;
    }
    button {
        border: none;
        font-size: 12px;
        background-color: inherit;
        border-radius: 15px;
    }
    button:hover {
        cursor: pointer;
        background-color: rgb(5, 253, 253, 0.7);
        font-weight: bold;
        color: #000000;;
    }
    .appList {
        margin: 2px auto;
        display: inline-block;
        margin-left: 2%;
    }
    #logo {
        margin: 2px;
        margin-left: 2%;
        font-weight: bold;
        color: #070606;
    }
    #logo:hover {
        background-color: inherit;
        cursor: pointer;
        color: #780ef1;
    }
    #exit {
        margin: 2px auto;
        margin-right: 2%;
        margin-left: auto;
    }
    #exit:hover {
        background-color: rgba(187, 8, 8, 0.7);
        color: #FFFFFF;
        font-weight: normal;
        font-weight: bold;
    }

</style>