<script>
    import {Route, Router} from "svelte-navigator";
    import Students from "./students/Students.svelte";
    import Teachers from "./teachers/Teachers.svelte";
</script>

<div class="wrapper">
    <Router primary={false}>
        <Route path="students">
            <Students/>
        </Route>
        <Route path="students">
            <Teachers/>
        </Route>
    </Router>
</div>

<style>
    .wrapper {
        position: relative;
        overflow-y: auto;
        max-height: 89vh;
        padding-top: 15px;
    }

    .wrapper::-webkit-scrollbar {
        width: 8px;
    }

    .wrapper::-webkit-scrollbar-thumb {
        background-color: black;
    }

    .wrapper::-webkit-scrollbar-track {
        -webkit-box-shadow: 5px 5px 5px -5px rgba(34, 60, 80, 0.2) inset;
        background-color: #f9f9fd;
    }
</style>
