import Link from "next/link";
import { Typography } from "@/components/Typography";
import styles from "./style.module.css";

export function Footer() {
  return (
    <footer className={styles.footer}>
      <Typography size="small" className={styles.copyright}>
        © Photo Share. All rights reserved.
      </Typography>
      <ul className={styles.list}>
        <li>
          {/* TODO: URL変更 */}
          <Link href="/">プライバシー・ポリシー</Link>
        </li>
        <li>
          {/* TODO: URL変更 */}
          <Link href="/">運営企業</Link>
        </li>
      </ul>
    </footer>
  );
}
