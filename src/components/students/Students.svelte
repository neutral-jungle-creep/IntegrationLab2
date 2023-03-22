<script>
    import {link} from "../../stores/stores.js";
    import {onMount} from "svelte";

    const courses = [1, 2, 3, 4, 5];
    let faculties = "";
    let groups = "";

    let selectFaculty = "15";
    let selectCourse = "1";
    let selectGroups = "";

    async function getFaculties() {
        console.log(link + "faculties");
        try {
            const response = await fetch(link + "faculties", {
                method: "GET",
            });

            if (response.status === 200) {
                faculties = await response.json();
            }
        } catch (err) {
            console.log(err);
        }

    }

    async function getGroups() {
        console.log(link + "groups");
        try {
            const response = await fetch(link + "groups", {
                method: "GET",
            });

            if (response.status === 200) {
                groups = await response.json();
            }
        } catch (err) {
            console.log(err);
        }

    }

    onMount(() => {
        getFaculties()
        getGroups()
    });

    async function setFacultyFilter() {
        selectFaculty = document.getElementById("faculty").value;
        console.log("факультет", selectFaculty);
    }

    async function setCourseFilter() {
        selectCourse = document.getElementById("course").value;
        console.log("курс", selectCourse);
    }

    async function setGroupFilter() {
        selectGroups = document.getElementById("group").value;
        console.log("группа", selectGroups);
    }

</script>

<div class="container">
    <div class="row">
        <h4>Расписание занятий</h4>
        <div class="col-md-5 col-sm-12 col-xs-12">
            <p>Выберете факультет</p>
            <select class="form-select" on:click={setFacultyFilter} id="faculty" required>
                {#each faculties as faculty}
                    {#if !faculty.name.startsWith("УДАЛЕН")}
                        <option value={faculty.facultyOid}>{faculty.name}</option>
                    {/if}
                {/each}
            </select>
        </div>
        <div class="col-md-2 col-sm-12 col-xs-12">
            <p>Выберете курс</p>
            <select class="form-select" on:click={setCourseFilter} id="course" required>
                {#each courses as course}
                    <option value={course}>{course}</option>
                {/each}
            </select>
        </div>
        <div class="col-md-5 col-sm-12 col-xs-12">
            <p>Выберете группу</p>
            <select class="form-select" on:click={setGroupFilter} id="group" required>
                {#each groups as group}
                    {#if group.facultyOid == selectFaculty && group.course == selectCourse}
                        <option value={group.groupOid}>{group.number + " " + group.speciality}</option>
                    {/if}
                {/each}
            </select>
        </div>
    </div>
    <div class="row">
        <div class="col-md-12 col-sm-12 col-xs-12">
            <p>расписание тут</p>
        </div>
    </div>
</div>

