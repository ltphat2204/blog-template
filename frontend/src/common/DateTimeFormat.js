export default function DateTimeFormat(date, setting) {
    const data = new Date(date);
    return new Intl.DateTimeFormat("en-US", setting).format(data);
}