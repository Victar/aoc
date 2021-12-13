package adventofcode.year2021;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Set;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day13 extends BaseTest {

	int sizeX = Integer.MIN_VALUE;
	int sizeY = Integer.MIN_VALUE;

	@Test public void runBoth() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day13/input.txt");
		Set<Point> points = new HashSet<>();
		final ArrayList<String> instructions = new ArrayList<String>();
		for (final String input : data) {
			if (StringUtils.contains(input, ",")) {
				final String[] arr = StringUtils.split(input, ",");
				final int curX = Integer.parseInt(arr[0]);
				final int curY = Integer.parseInt(arr[1]);
				this.sizeX = Math.max(this.sizeX, curX);
				this.sizeY = Math.max(this.sizeY, curY);
				final Point currentPoint = new Point(curX, curY);
				points.add(currentPoint);
			} else if (StringUtils.contains(input, "fold")) {
				instructions.add(input);
			}
		}
		boolean silverSolved = false;
		for (final String instruction : instructions) {
			points = foldPoints(points, instruction);
			if (!silverSolved) {
				System.out.println(points.size());
				silverSolved = true;
			}
		}
		drawPoint(points);
	}

	public Set<Point> foldPoints(final Set<Point> points, final String instruction) {
		final boolean isY = instruction.contains("fold along y=");
		final int direction = Integer.parseInt(StringUtils.split(instruction, "=")[1]);
		final Set<Point> result = new HashSet<>();
		for (final Point point : points) {
			if (isY) {
				final int x = point.x;
				final int y = point.y <= direction ? point.y : direction - (point.y - direction);
				result.add(new Point(x, y));
			} else {
				final int y = point.y;
				final int x = point.x <= direction ? point.x : direction - (point.x - direction);
				result.add(new Point(x, y));
			}
		}
		if (isY) {
			this.sizeY = direction;
		} else {
			this.sizeX = direction;
		}
		return result;
	}

	public void drawPoint(final Set<Point> points) {
		for (int i = 0; i <= this.sizeY; i++) {
			for (int j = 0; j <= this.sizeX; j++) {
				final String draw = points.contains(new Point(j, i)) ? "#" : ".";
				System.out.print(draw);
			}
			System.out.println("   y->" + i);
		}
	}

	@Data class Point {

		int x;
		int y;

		public Point(final int x, final int y) {
			this.x = x;
			this.y = y;
		}
	}

}
