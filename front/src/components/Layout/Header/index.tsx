import Link from "next/link";
import { Avatar } from "@/components/Avatar";
import styles from "./style.module.css";

export function Header() {
  return (
    <header className={styles.header}>
      <div className={styles.link}>
        <Link href="/">
          <p className={styles.siteName}>Photo Share</p>
        </Link>
        <div className={styles.avatar}>
          <Avatar />
          <div className={styles.action}>
            {/* Buttonにする */}
            ログイン
          </div>
        </div>
      </div>
    </header>
  );
}
