package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

public class Day14 extends BaseTest {

	public static final int DAY = 14;

	public static int MAX_Y = 0;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}


	@Test public void runSilver() throws Exception {
		runAny( false);

	}


	@Test public void runGold() throws Exception {
		runAny( true);
	}

	public void runAny(boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		List<Point> listPoint = new ArrayList<>();
		MAX_Y = 0;
		for (final String input : data) {
			String[] borders = input.split(" -> ");
			for (int i = 1; i < borders.length; i++) {
				listPoint.addAll(Point.getPoints(borders[i - 1], borders[i]));
			}
		}
		if (isGold){
			listPoint.addAll(Point.getPoints("-1000," + (MAX_Y + 1), "1500," + (MAX_Y + 1)));
		}
		int before = listPoint.size();
		boolean canAdd = true;
		int count = 0;
		while (canAdd) {
			count++;
			if (count>100000){ // Todo fix bug with enfless loop
				break;
			}
			Point start = Point.of(500, 0);
			while (start != null) {
				if (start.y >= MAX_Y) {
					canAdd = false;
					break;
				}
				start = addPoint(listPoint, start);
			}
		}
		System.out.println("size: " + (listPoint.size() - before));
	}

	public Point addPoint(List<Point> listPoint, Point current) {
		if (!listPoint.contains(Point.of(current.x, current.y + 1))) {
			return Point.of(current.x, current.y + 1);
		}
		if (!listPoint.contains(Point.of(current.x - 1, current.y + 1))) {
			return Point.of(current.x - 1, current.y + 1);
		}
		if (!listPoint.contains(Point.of(current.x + 1, current.y + 1))) {
			return Point.of(current.x + 1, current.y + 1);
		}
		if (!listPoint.contains(current)) {
			listPoint.add(current);
			return null;
		}
		return null;
	}

	@Data static class Point {

		int x;
		int y;

		public static Point of(int x, int y) {
			Point p = new Point();
			p.x = x;
			p.y = y;
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
					wall.add(Point.of(x, y));
				}
			}
			return wall;
		}
	}

}
