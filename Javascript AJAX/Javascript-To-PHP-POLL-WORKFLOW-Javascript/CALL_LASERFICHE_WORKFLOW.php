<?php


    $inputParameter = $_POST["inputParameter"];
    $uuid = $_POST["uuid"];

    if(isset($inputParameter) && isset($uuid)) {

        $handle = curl_init();


        $url = "http://WORKFLOW_URL_HERE";
    
        $data = ["ParameterCollection" => [
            
                                            ["Name" => "inputParameter",
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
            CURLOPT_USERPWD  => "DOMAIN\\XXXXXX:XXXXXX",
            CURLOPT_RETURNTRANSFER => true
        )
    
      );
    
        $output = curl_exec($handle);

        $statusCode = curl_getinfo($handle, CURLINFO_HTTP_CODE);


        if($statusCode == 200) {
            touch("WORKFLOW_TEXT_FILES/".$uuid."_STARTED.txt");
        }

        echo $statusCode;

        curl_close($handle);
    
    }

?>