<script>
    import {link} from "../../stores/stores.js";
    import {onMount} from "svelte";

    const courses = [1, 2, 3, 4, 5];
    let faculties = "";
    let groups = "";

    let selectFaculty = "15";
    let selectCourse = "1";
    let selectGroups = "9468";

    let lessons = "";

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

    async function getLessons() {
        let start = document.getElementById("month").value;
        let end = Number(start) + 1 + ".1";

        if (start == 12) {
            end = 12 + ".31"
        }

        const params = {
            groupOid: selectGroups,
            fromdate: "2018." + start + ".1",
            todate: "2018." + end
        }

        try {
            const response = await fetch(link + "lessons?" + (new URLSearchParams(params)), {
                method: "GET",
            });

            if (response.status === 200) {
                lessons = await response.json();
                console.log(lessons)
            }
        } catch (err) {
            console.log(err);
        }

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
    <div class="row mt-4">
        <div class="col-md-4 col-sm-12 col-xs-12 d-grid">
            <select class="form-select" on:click={getLessons} id="month" required>
                <option value="1">январь</option>
                <option value="2">февраль</option>
                <option value="3">март</option>
                <option value="4">апрель</option>
                <option value="5">май</option>
                <option value="6">июнь</option>
                <option value="7">июль</option>
                <option value="8">август</option>
                <option value="9">сентябрь</option>
                <option value="10">октябрь</option>
                <option value="11">ноябрь</option>
                <option value="12">декабрь</option>
            </select></div>
        <div class="row mt-4">

            <div class="col-md-12 col-sm-12 col-xs-12">
                <table class="table col-12 mt-1">
                    <tr class="text-uppercase fw-bold">
                        <th class="col-1">Дата и время</th>
                        <th class="col-2">Здание</th>
                        <th class="col-3">Аудитория</th>
                        <th class="col-4">Предмет</th>
                        <th class="col-5">Преподаватель</th>
                    </tr>
                    {#if lessons.length == 0}
                        <br>
                        <tr>
                        <td class="label">Занятий нет</td>
                        </tr>
                    {:else}
                        {#each lessons as lesson}
                            <br>
                            <tr>
                                <td>{lesson.date}<br>{lesson.beginLesson + " - " + lesson.endLesson}</td>
                                <td>{lesson.building}</td>
                                <td>{lesson.auditorium}</td>
                                <td>{lesson.discipline}</td>
                                <td>{lesson.lecturer}</td>
                            </tr>
                        {/each}
                    {/if}

                </table>
            </div>
        </div>
    </div>
</div>

<style>
    table {
        table-layout: fixed;
        text-align: center;
    }

    .label {
        text-align: center;
        color: crimson;
        font-size: 30px;
    }

    .col-1 {
        width: 10%;
    }

    .col-2 {
        width: 20%;
    }

    .col-3 {
        width: 20%;
    }

    .col-4 {
        width: 30%;
    }

    .col-5 {
        width: 20%;
    }
</style>