"use client";

import type { ReactNode } from "react";
import clsx from "clsx";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Icon } from "@/components/Icon";
import { PhotoCreateModalContainer } from "../PhotoCreateModalContainer";
import styles from "./style.module.css";

function renderLink(
  flag: boolean,
  renderer: (attr?: { "aria-current": "page" }) => ReactNode
) {
  return renderer(flag ? { "aria-current": "page" as const } : undefined);
}

export function Navigation() {
  const currentPathname = usePathname();

  return (
    <nav className={styles.nav}>
      <ul className={styles.list}>
        <li className={styles.listitem}>
          {renderLink(currentPathname === "/", (attr) => (
            <Link href="/" className={styles.navigationLinkClassName} {...attr}>
              <Icon type="home" color={Boolean(attr) ? "orange" : "black"} />
              home
            </Link>
          ))}
        </li>
        <li className={styles.listitem}>
          {renderLink(currentPathname === "/profile", (attr) => (
            <Link href="/" className={styles.navigationLinkClassName} {...attr}>
              <Icon type="user" color={Boolean(attr) ? "orange" : "black"} />
              profile
            </Link>
          ))}
        </li>
        <li className={styles.listitem}>
          <PhotoCreateModalContainer>
            <span
              className={clsx(
                styles.listitemChild,
                styles.navigationLinkClassName
              )}
            >
              <Icon type="camera" />
              post
            </span>
          </PhotoCreateModalContainer>
        </li>
      </ul>
    </nav>
  );
}
