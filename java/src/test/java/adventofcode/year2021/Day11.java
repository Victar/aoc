package adventofcode.year2021;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day11 extends BaseTest {

	int resultSilver;
	int resultGold;

	@Test public void runSilver() throws Exception {
		run(false);
		System.out.println(this.resultSilver);
	}

	@Test public void runGold() throws Exception {
		run(true);
		System.out.println(this.resultGold);
	}

	public void run(final boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day11/input.txt");
		final int SIZE = data.size();
		ArrayList<Point> points = new ArrayList<>();
		for (int i = 0; i < SIZE; i++) {
			for (int j = 0; j < SIZE; j++) {
				points.add(new Point(i, j, Integer.parseInt("" + data.get(i).charAt(j))));
			}
		}
		for (int i = 0; isGold ? this.resultGold == 0 : i < 100; i++) {
			points = doRound(points, SIZE, i + 1);
		}
	}

	public ArrayList<Point> doRound(final ArrayList<Point> points, final int size, final int current) {
		final ArrayList<Point> result = new ArrayList<>(points);
		for (final Point point : result) {
			point.setVisited(false);
			point.inc();
		}
		for (final Point point : result) {
			flush(result, point, size);
		}
		int count = 0;
		for (final Point point : result) {
			if (point.isVisited()) {
				point.setValue(0);
				count++;
			}
		}
		if (count == size * size) {
			this.resultGold = current;
		}
		return result;
	}

	public void flush(final ArrayList<Point> result, final Point point, final int size) {
		final int current = point.getValue();
		if (current > 9 && !point.isVisited()) {
			for (int i = point.getX() - 1; i <= point.getX() + 1; i++) {
				for (int j = point.getY() - 1; j <= point.getY() + 1; j++) {
					if (i < size && j < size && i >= 0 && j >= 0) {
						result.get(i * size + j).inc();
					}
				}
			}
			point.setVisited(true);
			this.resultSilver++;
			for (final Point p : result) {
				if (p.getValue() > 9) {
					flush(result, p, size);
				}
			}
		}
	}

	@Data public static class Point {

		private int x;
		private int y;
		private int value;
		private boolean visited;

		public Point(final int x, final int y, final int value) {
			this.x = x;
			this.y = y;
			this.value = value;
		}

		public void inc() {
			this.value++;
		}
	}

}
