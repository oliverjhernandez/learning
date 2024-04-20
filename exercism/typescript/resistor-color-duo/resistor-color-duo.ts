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

export function decodedValue(resistors: ResistorColor[]): number {
  const firstValue = RESISTORS[resistors[0]];
  const secondValue = RESISTORS[resistors[1]];
  return firstValue * 10 + secondValue;
}
