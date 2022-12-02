package adventofcode.year2021;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;
import lombok.RequiredArgsConstructor;

public class Day5 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day5/input.txt");
		System.out.println(countPoints(data, false));
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day5/input.txt");
		System.out.println(countPoints(data, true));

	}

	public int countPoints(final ArrayList<String> data, final boolean countGold) {
		final Map<Point, Integer> results = new HashMap<>();
		for (final String input : data) {
			final boolean skip = true;
			final String[] arr1 = StringUtils.split(input, " -> ");
			final String[] arr11 = StringUtils.split(arr1[0], ",");
			final int x1 = Integer.parseInt(arr11[0]);
			final int y1 = Integer.parseInt(arr11[1]);
			final String[] arr12 = StringUtils.split(arr1[1], ",");
			final int x2 = Integer.parseInt(arr12[0]);
			final int y2 = Integer.parseInt(arr12[1]);
			if (x1 == x2) {
				for (int j = Math.min(y1, y2); j <= Math.max(y1, y2); j++) {
					final Point currentPoint = new Point(x1, j);
					if (results.containsKey(currentPoint)) {
						results.put(currentPoint, results.get(currentPoint) + 1);
					} else {
						results.put(currentPoint, 1);
					}
				}
			}
			if (y1 == y2) {
				for (int j = Math.min(x1, x2); j <= Math.max(x1, x2); j++) {
					final Point currentPoint = new Point(j, y1);
					if (results.containsKey(currentPoint)) {
						results.put(currentPoint, results.get(currentPoint) + 1);
					} else {
						results.put(currentPoint, 1);
					}
				}
			}
			if (countGold) {
				if (Math.abs(x1 - x2) == Math.abs(y1 - y2)) {
					final int size = Math.abs(x1 - x2);
					final int directionX = x1 >= x2 ? -1 : 1;
					final int directionY = y1 >= y2 ? -1 : 1;
					for (int l = 0; l <= size; l++) {
						final Point currentPoint = new Point(x1 + l * directionX, y1 + l * directionY);
						if (results.containsKey(currentPoint)) {
							results.put(currentPoint, results.get(currentPoint) + 1);
						} else {
							results.put(currentPoint, 1);
						}
					}
				}
			}
		}
		int count = 0;
		for (final Map.Entry<Point, Integer> point : results.entrySet()) {
			if (point.getValue() > 1) {
				count++;
			}
		}
		return count;
	}

	@Data @RequiredArgsConstructor(staticName = "of") public static class Point {

		private int x;
		private int y;

		public Point(final int x, final int y) {
			this.x = x;
			this.y = y;
		}
	}

}
