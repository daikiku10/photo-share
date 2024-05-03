import type { ReactNode } from "react";
import React from "react";
import styles from "./style.module.css";

export function Root({ children }: { children: ReactNode }) {
  return <div className={styles.root}>{children}</div>;
}
