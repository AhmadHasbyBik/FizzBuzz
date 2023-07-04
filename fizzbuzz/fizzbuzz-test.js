console.log("SOAL!\nBuat loop yang mencetak angka 1-100, JIKA kelipatan 3 maka outputnya 'Fizz', \nJIKA kelipatan 5 outputnya 'Buzz' dan JIKA kelipatan 3 dan 5 keluarkan 'FizzBuzz'\nJAWAB!");
for (let i = 1; i <= 100; i++) {
  if (i % 3 === 0 && i % 5 === 0) {
    console.log("FizzBuzz");
  } else if (i % 3 === 0) {
    console.log("Fizz");
  } else if (i % 5 === 0) {
    console.log("Buzz");
  } else {
    console.log(i);
  }
}
