import type { ComponentPropsWithRef } from "react";
import React, { forwardRef } from "react";
import clsx from "clsx";
import styles from "./style.module.css";

type Props = ComponentPropsWithRef<"main">;
export const Main = forwardRef<HTMLElement, Props>(function MainBase(
  { className, ...props },
  ref
) {
  return <main {...props} ref={ref} className={clsx(styles.main, className)} />;
});
