<script>
    import {Router} from "svelte-navigator";
</script>


<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container">
        <a class="navbar-brand" href="/">РАСПИСАНИЕ</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
                aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>

        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <Router>
                <li class="nav-item">
                    <a class="nav-link" href="students">Студенту</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="teachers">Преподавателю</a>
                </li>
            </Router>
        </ul>
    </div>
</nav>


<style></style>