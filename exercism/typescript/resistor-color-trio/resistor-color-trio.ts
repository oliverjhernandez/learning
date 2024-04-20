type ResistorColor = keyof typeof RESISTORS;

const RESISTORS = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
};

export function decodedResistorValue(colors: ResistorColor[]): string {
  const firstValue = RESISTORS[colors[0]];
  const secondValue = RESISTORS[colors[1]];
  const thirdValue = RESISTORS[colors[2]];

  return (
    String(Number(`${firstValue}${secondValue}`) + 10 * thirdValue) + " ohms"
  );
}
