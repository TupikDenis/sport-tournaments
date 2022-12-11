$(document).ready(function() {
    $("input[name='system']").change(function(){
        if($("input[name='system']:checked").val() === "robin"){
            var html = `<div class="card m-5" id="robin_round">
                                <h5 class="card-header">Robin round</h5>
                                <div class="card-body">
                                    <div class="form-group mt-3">
                                        <label for="team_number_robin"><b>Number of teams (4-10)</b></label>
                                        <input type="number" class="form-control" name="team_number_robin" id="team_number_robin" min="4" value="4" max="10">
                                    </div>
                                    
                                    <div class="form-group mt-3">
                                        <label for="number_robin"><b>Number of robins (1-4)</b></label>
                                        <input type="number" class="form-control" name="number_robin" id="number_robin" min="1" value="1" max="4">
                                    </div>

                                    <h5 class="mt-5">Teams</h5>
                                    <div class="teams" id="team_robin">
                                        <div class="form-group mt-1">
                                            <label for="team_name1_robin"><b>Team #1</b></label>
                                            <input type="text" class="form-control" name="team_name1_robin" id="team_name1_robin" placeholder="Team #1"></div>
                                        <div class="form-group mt-1">
                                            <label for="team_name2_robin"><b>Team #2</b></label>
                                            <input type="text" class="form-control" name="team_name2_robin" id="team_name2_robin" placeholder="Team #2"></div>
                                       <div class="form-group mt-1">
                                             <label for="team_name3_robin"><b>Team #3</b></label>
                                             <input type="text" class="form-control" name="team_name3_robin" id="team_name3_robin" placeholder="Team #3">
                                       </div>
                                       <div class="form-group mt-1">
                                             <label for="team_name4_robin"><b>Team #4</b></label>
                                             <input type="text" class="form-control" name="team_name4_robin" id="team_name4_robin" placeholder="Team #4">
                                       </div>
                                    </div>

                                    <div class="form-group mt-5">
                                        <div class="form-check">
                                            <input class="form-check-input" type="checkbox" name="mixed_team_robin" value="mixed_team_robin" id="mixed_team_robin">
                                            <label class="form-check-label" for="mixed_team_robin"><b>Mixed Team</b></label>
                                        </div>
                                    </div>
                                </div>
                            </div>`
        }

        if($("input[name='system']:checked").val() === "knockout"){
            var html = `<div class="card m-5" id="knockout_round">
                            <h5 class="card-header">Knockout system</h5>
                            <div class="card-body">
                                <div class="form-group mt-3">
                                    <div class="form-group mt-3">
                                        <label for="team_number_knockout"><b>Number of teams (4-16)*</b></label>
                                        <input type="number" class="form-control" name="team_number_knockout" id="team_number_knockout" min="4" value="4" max="16">
                                    </div>
                                </div>
                
                                <h5 class="mt-5">Teams</h5>
                                <div class="team_knockout" id="team_knockout">
                                        <div class="form-group mt-1">
                                            <label for="team_name1_knockout"><b>Team #1</b></label>
                                            <input type="text" class="form-control" name="team_name1_knockout" id="team_name1_knockout" placeholder="Team #1"></div>
                                        <div class="form-group mt-1">
                                            <label for="team_name2_knockout"><b>Team #2</b></label>
                                            <input type="text" class="form-control" name="team_name2_knockout" id="team_name2_knockout" placeholder="Team #2"></div>
                                        <div class="form-group mt-1">
                                             <label for="team_name3_knockout"><b>Team #3</b></label>
                                             <input type="text" class="form-control" name="team_name3_knockout" id="team_name3_knockout" placeholder="Team #3">
                                       </div>
                                       <div class="form-group mt-1">
                                             <label for="team_name4_knockout"><b>Team #4</b></label>
                                             <input type="text" class="form-control" name="team_name4_knockout" id="team_name4_knockout" placeholder="Team #4">
                                       </div>
                                </div>
                
                                <div class="form-group mt-3">
                                    <div class="form-check">
                                        <input class="form-check-input" type="checkbox" value="mixed_team_knockout" id="mixed_team_knockout">
                                        <label class="form-check-label" for="mixed_team_knockout"><b>Mixed Team</b></label>
                                    </div>
                                </div>
                            </div>
                        </div>`
        }


        $("#results").html(html).val();
    });

    $(document).on('input', '#team_number_robin', function(){
        var numberTeam = $("#team_number_robin").val();
        var container = ''
        for(let i=0; i < numberTeam; i++){
            var number = i+1
            container += '<div class="form-group mt-1">'
            container += '<label for="team_name'+ number +'_robin">'
            container += '<b>Team #'+ number +'</b></label>'
            container += '<input type="text" class="form-control" name="team_name'+ number +'_robin" id="team_name'+ number +'_robin" placeholder="Team #'+ number +'"></div>'
        }

        $('#team_robin').html(container);
    })

    $(document).on('input', '#team_number_knockout', function(){
        var numberTeam = $("#team_number_knockout").val();
        var container = ''
        for(let i=0; i < numberTeam; i++){
            var number = i+1
            container += '<div class="form-group mt-1">'
            container += '<label for="team_name'+ number +'_knockout">'
            container += '<b>Team #'+ number +'</b></label>'
            container += '<input type="text" class="form-control" name="team_name'+ number +'_knockout" id="team_name'+ number +'_knockout" placeholder="Team #'+ number +'"></div>'
        }

        $('#team_knockout').html(container);
    })
});