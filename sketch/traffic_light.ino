const int RED = 8;
const int YELLOW = 7;
const int GREEN = 6;

boolean YELLOW_BLINKING = false;

void offAll() {
  digitalWrite(RED, LOW);
  digitalWrite(YELLOW, LOW);
  digitalWrite(GREEN, LOW);
}

void off(int led) {
  digitalWrite(led, LOW);
}

void on(int led) {
  digitalWrite(led, HIGH);
}

void setup() {
  // initialize the digital pin as an output.
  pinMode(RED, OUTPUT);
  pinMode(YELLOW, OUTPUT);
  pinMode(GREEN, OUTPUT);
  // initialize serial:
  Serial.begin(9600);
}

void loop() {
    if(YELLOW_BLINKING) {
      on(YELLOW);
      delay(500);
      off(YELLOW);
    }
}

void serialEvent() {
  while (Serial.available()) {
    // get the new byte:
    char inChar = (char)Serial.read();
    // add it to the inputString:
    switch (inChar) {
      case 'R':
        offAll();
        on(RED);
        break;
      case 'Y':
        offAll();
        on(YELLOW);
        break;
      case 'y':
        offAll();
        YELLOW_BLINKING = YELLOW_BLINKING ? false : true;
      case 'G':
        offAll();
        on(GREEN);
        break;
      case 'O':
        offAll();
        break;
    }
  }
}
