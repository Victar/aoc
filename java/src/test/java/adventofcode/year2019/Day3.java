package adventofcode.year2019;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Set;
import java.util.stream.IntStream;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;
import lombok.RequiredArgsConstructor;

public class Day3 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day3/input.txt");
		final Set<Point> set0 = stringToRoute(data.get(0));
		final Set<Point> set1 = stringToRoute(data.get(1));
		System.out.println(set0);
		System.out.println(set1);
		set0.retainAll(set1); // s1 now contains only elements in both sets
		System.out.println(set0);
		int minDistance = Integer.MAX_VALUE;
		final ArrayList<Point> arr = new ArrayList(set0);
		for (int i = 0; i < arr.size(); i++) {
			final Point point1 = arr.get(i);
			final int currentDistance = Math.abs(point1.x) + Math.abs(point1.y);
			minDistance = Math.min(minDistance, currentDistance);
		}
		System.out.println(minDistance);

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day3/input.txt");
		final ArrayList<Point> list0 = stringToRouteGold(data.get(0));
		final ArrayList<Point> list1 = stringToRouteGold(data.get(1));
		final Set<Point> set0 = new HashSet<>(list0);
		final Set<Point> set1 = new HashSet<>(list1);
		set0.retainAll(set1); // s1 now contains only elements in both sets
		System.out.println(set0);
		int minDistance = Integer.MAX_VALUE;
		final ArrayList<Point> arr = new ArrayList(set0);
		for (int i = 1; i < arr.size(); i++) {
			final Point point1 = arr.get(i);
			final int route1 = IntStream.range(0, list0.size()).filter(userInd -> list0.get(userInd).equals(point1)).findFirst().getAsInt();
			final int route2 = IntStream.range(0, list1.size()).filter(userInd -> list1.get(userInd).equals(point1)).findFirst().getAsInt();
			final int currentDistance = route1 + route2;
			minDistance = Math.min(minDistance, currentDistance);
		}

		System.out.println(minDistance);

	}

	ArrayList<Point> stringToRouteGold(final String data) {
		final String[] arr = data.split(",");
		final ArrayList<Point> result = new ArrayList<>();
		result.add(new Point(0, 0));
		int currentX = 0;
		int currentY = 0;
		for (final String item : arr) {
			final String direction = item.substring(0, 1);
			final int distance = Integer.parseInt(item.substring(1));
			for (int i = 0; i < distance; i++) {
				if ("R".equals(direction)) {
					currentX++;
				}
				if ("L".equals(direction)) {
					currentX--;
				}
				if ("U".equals(direction)) {
					currentY++;
				}
				if ("D".equals(direction)) {
					currentY--;
				}
				result.add(new Point(currentX, currentY));
			}
		}
		return result;
	}

	Set<Point> stringToRoute(final String data) {
		final String[] arr = data.split(",");
		final Set<Point> result = new HashSet<>();
		int currentX = 0;
		int currentY = 0;
		for (final String item : arr) {
			final String direction = item.substring(0, 1);
			final int distance = Integer.parseInt(item.substring(1));
			for (int i = 0; i < distance; i++) {
				if ("R".equals(direction)) {
					currentX++;
				}
				if ("L".equals(direction)) {
					currentX--;
				}
				if ("U".equals(direction)) {
					currentY++;
				}
				if ("D".equals(direction)) {
					currentY--;
				}
				result.add(new Point(currentX, currentY));
			}
		}
		return result;
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
