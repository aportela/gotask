export class IDate {
  msTimestamp: number | null;
  date: Date | null;

  constructor(msTimestamp: number | null) {
    this.msTimestamp = msTimestamp;
    this.date = msTimestamp !== null ? new Date(msTimestamp) : null;
  }

  toLocaleString = () => {
    return this.date?.toLocaleString();
  };
}
