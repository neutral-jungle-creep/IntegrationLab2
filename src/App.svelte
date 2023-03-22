<script>
    import Header from "./components/Header.svelte"
    import Workspace from "./components/Workspace.svelte";
</script>

<svelte:head>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css" />
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.9.1/font/bootstrap-icons.css">
</svelte:head>

<main>

    <div class="app-container">
        <Header/>
        <Workspace/>
    </div>

</main>

<style>

</style>
