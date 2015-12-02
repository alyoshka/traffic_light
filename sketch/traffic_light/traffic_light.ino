const int RED = 8;
const int YELLOW = 7;
const int GREEN = 6;

boolean YELLOW_BLINKING;

void offAll() {
  YELLOW_BLINKING = false;
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
  YELLOW_BLINKING = false;
  // initialize the digital pin as an output.
  pinMode(RED, OUTPUT);
  pinMode(YELLOW, OUTPUT);
  pinMode(GREEN, OUTPUT);
  // initialize serial:
  Serial.begin(9600);
}

void loop() {
    if(YELLOW_BLINKING == true) {
      on(YELLOW);
      delay(500);
      off(YELLOW);
      delay(500);
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
        if(YELLOW_BLINKING == true) {
          YELLOW_BLINKING = false;
        } else {
          YELLOW_BLINKING = true;
        }
        break;
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
