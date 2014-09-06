// Declare some global variables
JSONArray states;
float X; // position of cart
float Theta1; // long pole angle
float Theta2; // short pole angle

void setup(){
  // Load the JSON file as a JSON object
  states = loadJSONArray("../../processingjs/json/states.json");
  JSONObject state = states.getJSONObject(0);
  //println(states);
  size(480, 200);
  noStroke();
  smooth();
  noLoop();
}

void draw()
{
  background(102);
  drawCart();
  drawPoles();
  println("{X : " + X + " (cm), " + "Theta1 : " + Theta1 + ", "  + "Theta2 : " + Theta2 + "}");
}

void drawPoles()
{
  stroke(255);
  strokeWeight(3);
  // short pole
  pushMatrix();
  // move the origin to the pivot point
  translate(240+X, 150);
  rectMode(CORNER);
  // then pivot the grid - as this is what moves
  rotate(Theta2 + radians(180));
  fill(0.2);
  // and draw the square at the new origin
  rect(0, 0, 1, 20);
  popMatrix();

  // long pole
  pushMatrix();
  // move the origin to the pivot point
  translate(240+X, 150);
  rectMode(CORNER);
  // then pivot the grid - as this is what moves
  rotate(Theta1 + radians(180));
  fill(0.2);
  // and draw the square at the new origin
  rect(0, 0, 1, 100);
  popMatrix();
}

void drawCart(){
  fill(0.2);
  rectMode(CORNER);
  rect(190+X, 150, 100, 50);
}
