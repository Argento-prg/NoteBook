<script>
    import { store } from './storeApp'
</script>

<div class="home">
    Добро пожаловать в NoteBook, <span>{store.getName()}</span>!<br />
    Выберите приложение, с которым хотите начать работу, щёлкнув по нему вверху страницы.
</div>

<style>
    .home {
        text-align: center;
        margin-top: 10%;
        font-size: 20px;
    }
    span {
        font-weight: bold;
        font-style: italic;
    }
</style>