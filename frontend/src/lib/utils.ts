import type { TLog } from "./state";

const HEADER_SEP_WIDTH = 63;

export function logsToString(logs: TLog[]): string {
  const col_sep = "\t|\t";
  const header = "Timestamp" + col_sep + "Level" + col_sep + "Message" + "\n";
  const header_sep = "─".repeat(HEADER_SEP_WIDTH) + "\n";
  const log_lines: string[] = logs.map((log: TLog): string => {
    const timestamp = new Date(log.timestamp * 1000)
      .toTimeString()
      .substr(0, 8);
    return timestamp + col_sep + log.level + col_sep + log.message;
  });
  return header + header_sep + log_lines.join("\n");
}
