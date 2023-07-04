<?php
echo "SOAL!\nBuat loop yang mencetak angka 1-100, JIKA kelipatan 3 maka outputnya 'Fizz', \nJIKA kelipatan 5 outputnya 'Buzz' dan JIKA kelipatan 3 dan 5 keluarkan 'FizzBuzz'\nJAWAB!\n";
for ($i = 1; $i <= 100; $i++) {
  if ($i % 3 === 0 && $i % 5 === 0) {
    echo "FizzBuzz\n";
  } else if ($i % 3 === 0) {
    echo "Fizz\n";
  } else if ($i % 5 === 0) {
    echo "Buzz\n";
  } else {
    echo $i . "\n";
  }
}
?>