export class IDate {
  msTimestamp: number;
  date: Date;

  constructor(msTimestamp: number) {
    this.msTimestamp = msTimestamp;
    this.date = new Date(msTimestamp);
  }

  toLocaleString = () => {
    return this.date.toLocaleString();
  };
}
