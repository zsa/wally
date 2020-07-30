import React from "react";

const LogLine = ({ line }) => {
  const timestamp = new Date(line.timestamp * 1000).toTimeString().substr(0, 8);
  return (
    <div className="line">
      <div className="timestamp">{timestamp}</div>
      <div className={`level ${line.level}`}>
        <div className="message">{line.level}</div>
      </div>
      <div className="message">{line.message}</div>
    </div>
  );
};

export default ({ logs }) => {
  return (
    <div className="console">
      <div className="lines">
        { logs && logs.map((line, index) => (
          <LogLine key={index} line={line} />
        ))}
      </div>
    </div>
  );
};
