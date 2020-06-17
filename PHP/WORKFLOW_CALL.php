<?php

    $handle = curl_init();

    $url = "http://WORKFLOW_URL_HERE";

    $data = ["ParameterCollection" => [["Name" => "Message",
                                        "Value" => "Hello World!"],
                                       [
                                        "Name" => "Message Two",
                                        "Value" => "Hello World!"
                                       ]]];

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
        CURLOPT_USERPWD  => "domain\\ACCOUNT_HERE:XXXXXXXXX",
        CURLOPT_RETURNTRANSFER => true
    )

  );

    $output = curl_exec($handle);
	
    $statusCode = curl_getinfo($handle, CURLINFO_HTTP_CODE); // v=-A1_atkAQhY

    curl_close($handle);

	
	if($statusCode == 200) {
		echo "WORKFLOW STARTED SUCCESSFULLY";
	} else {
		echo "There was a problem starting workflow";
	}

?>
