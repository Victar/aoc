package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day14 extends BaseTest {

	public static final int DAY = 14;

	public static int MAX_Y = 0;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		runAny(false);
	}

	@Test public void runGold() throws Exception {
		runAny(true);
	}

	public void runAny(boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		List<Point> listPoint = new ArrayList<>();
		Set<String> setPointHashes = new HashSet<>();
		MAX_Y = 0;
		for (final String input : data) {
			String[] borders = input.split(" -> ");
			for (int i = 1; i < borders.length; i++) {
				listPoint.addAll(Point.getPoints(borders[i - 1], borders[i]));
			}
		}
		if (isGold) {
			listPoint.addAll(Point.getPoints("-1000," + (MAX_Y + 1), "1500," + (MAX_Y + 1)));
		}
		int before = listPoint.size();
		for (Point point : listPoint) {
			setPointHashes.add(point.getHash());
		}
		boolean canAdd = true;
		final Point START = Point.of(500, 0);
		while (canAdd) {
			Point start = START;
			while (start != null) {
				if (start.y >= MAX_Y) {
					canAdd = false;
					break;
				}
				start = addPoint(listPoint, setPointHashes, start);
			}
			if (listPoint.contains(START)) {
				canAdd = false;
			}
		}

		System.out.println((listPoint.size() - before));
		drawPoints(listPoint);
	}

	public Point addPoint(List<Point> listPoint, Set<String> setPointHashes, Point current) {
		if (!setPointHashes.contains(current.x + "," + (current.y + 1))) {
			return Point.of(current.x, current.y + 1);
		}
		if (!setPointHashes.contains((current.x - 1) + "," + (current.y + 1))) {
			return Point.of(current.x - 1, current.y + 1);
		}
		if (!setPointHashes.contains((current.x + 1) + "," + (current.y + 1))) {
			return Point.of(current.x + 1, current.y + 1);
		}
		if (!listPoint.contains(current)) {
			listPoint.add(current);
			setPointHashes.add(current.getHash());
			return null;
		}
		return null;
	}

	public void drawPoints(List<Point> listPoint) {
		Map<String, Point> map = new HashMap<>();
		int startX = Integer.MAX_VALUE;
		int startY = Integer.MAX_VALUE;
		int endX = Integer.MIN_VALUE;
		int endY = Integer.MIN_VALUE;
		for (Point point : listPoint) {
			if (!map.containsKey(point.getHash())) {
				map.put(point.getHash(), point);
			}
			if (!point.isWall()) {
				startX = Math.min(startX, point.getX() - 2);
				endX = Math.max(endX, point.getX() + 2);
				startY = Math.min(startY, point.getY() - 2);
				endY = Math.max(endY, point.getY() + 2);
			}
		}
		System.out.println("=Draw= (" + startX + "->" + endX + ")  (" + startY + "->" + endY + ")");
		for (int y = startY; y < endY; y++) {
			StringBuilder sb = new StringBuilder();
			for (int x = startX; x < endX; x++) {
				String cur = " ";
				Point p = map.get(x + "," + y);
				if (p != null) {
					cur = p.isWall ? "#" : "0";
				}
				sb.append(cur);
			}
			System.out.println(sb);
		}
		System.out.println("=End=\n\n");
	}

	@Data static class Point {

		int x;
		int y;
		boolean isWall = false;

		public static Point of(int x, int y) {
			Point p = new Point();
			p.x = x;
			p.y = y;
			return p;
		}

		public static Point ofWall(int x, int y) {
			Point p = new Point();
			p.x = x;
			p.y = y;
			p.isWall = true;
			return p;
		}

		public static List<Point> getPoints(String start, String end) {
			final List<Point> wall = new ArrayList();
			final String[] sArr = start.split(",");
			final String[] eArr = end.split(",");
			int sX = Math.min(Integer.parseInt(sArr[0]), Integer.parseInt(eArr[0]));
			int sY = Math.min(Integer.parseInt(sArr[1]), Integer.parseInt(eArr[1]));
			int eX = Math.max(Integer.parseInt(sArr[0]), Integer.parseInt(eArr[0]));
			int eY = Math.max(Integer.parseInt(sArr[1]), Integer.parseInt(eArr[1]));
			for (int x = sX; x <= eX; x++) {
				for (int y = sY; y <= eY; y++) {
					MAX_Y = Math.max(MAX_Y, y + 1);
					Point p = Point.of(x, y);
					p.setWall(true);
					wall.add(p);

				}
			}
			return wall;
		}

		public String getHash() {
			return x + "," + y;
		}

		@Override public boolean equals(Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;
			Point point = (Point) o;
			return x == point.x && y == point.y;
		}

		@Override public int hashCode() {
			return Objects.hash(x, y);
		}
	}

}
