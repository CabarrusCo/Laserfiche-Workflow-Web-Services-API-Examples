<?php

    $inputParameter = $_POST["inputParameter"];
    $uuid = $_POST["uuid"];

    if(isset($inputParameter) && isset($uuid)) {

        $handle = curl_init();


        $url = "http://WORKFLOW_URL_HERE";
    
        $data = ["ParameterCollection" => [
            
                                            [
                                                "Name" => "inputParameter",
                                                "Value" => $inputParameter
                                            ],
                                            [
                                                "Name" => "uuid",
                                                "Value" => $uuid
                                            ],
                                        ]
                ];
    
        $data = json_encode($data);
    
        curl_setopt_array($handle,
    
        array(
            CURLOPT_URL => $url,
            CURLOPT_POST => true,
    
            CURLOPT_HTTPHEADER => array(                                                                          
                'Content-Type: application/json',                                                                                
                'Content-Length: ' . strlen($data)
            ),
    
            CURLOPT_POSTFIELDS => $data,
            CURLOPT_HTTPAUTH => CURLAUTH_NTLM,
            CURLOPT_USERPWD  => "DOMAIN\\XXXXXXX:XXXXXXX",
            CURLOPT_RETURNTRANSFER => true
        )
    
      );
    
        $output = curl_exec($handle);

        $statusCode = curl_getinfo($handle, CURLINFO_HTTP_CODE);


        if($statusCode == 200) {

            $loopThroughDirectoryFlag = true;
            $loopCounter = 0;

            while($loopThroughDirectoryFlag == true) {

                $textFiles = scandir("WORKFLOW_TEXT_FILES");

                foreach($textFiles as $textFile) {

                    if (strpos($textFile, $uuid) !== false) {
                        $loopThroughDirectoryFlag = false;
                        $textFileStatus = explode("_", $textFile);
                        $textFileStatus = $textFileStatus[1];
                        $textFileStatus = str_replace(".txt", "", $textFileStatus);
                        echo "WORKFLOW FINISHED WITH STATUS OF ". $textFileStatus;
                        unlink("WORKFLOW_TEXT_FILES/".$textFile);
                    } else {
                        if($loopCounter <= 120) {
                            time.sleep(1);
                            $loopCounter = $loopCounter + 1;
                        } else {
                            $loopThroughDirectoryFlag = false;
                            echo "MAX ITERATION REACHED, NO FILE FOUND";
                        }

                    }
    
                }

            }



        } else {
            echo "WORKFLOW FAILED TO FIRE";
        }


        curl_close($handle);
    
    }

?>