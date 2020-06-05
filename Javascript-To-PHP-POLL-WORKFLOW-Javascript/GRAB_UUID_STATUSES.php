<?php

    $currentIdsInCirculation = $_POST["currentIdsInCirculation"];
    $currentIdsInCirculation = json_decode($currentIdsInCirculation);

    $statusArrayForReturn = array();

    foreach($currentIdsInCirculation as $currentID) {
        $uuidForFile = $currentID->uuid;
        $workflowsInCirculationFiles = scandir("WORKFLOW_TEXT_FILES");

        foreach($workflowsInCirculationFiles as $workflowTextFile) {

            if (strpos($workflowTextFile, $uuidForFile) !== false) {

                $textFileStatus = explode("_", $workflowTextFile);
                $textFileStatus = $textFileStatus[1];
                $textFileStatus = str_replace(".txt", "", $textFileStatus);


                array_push($statusArrayForReturn, ["uuid" => $uuidForFile, "status" => $textFileStatus]);

            }

        }
    }


    echo json_encode($statusArrayForReturn);
    
?>