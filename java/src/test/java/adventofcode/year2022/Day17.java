package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Day17 extends BaseTest {

	public static final int DAY = 17;

	public static final int X_SIZE = 7;
	static int cursor = 0;


	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test
	public void runBoth() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/sample.txt");
		String pattern = data.get(0);
		Board board = new Board();
		Map<Long, Long> sizeAdded = new HashMap<>();
		StringBuilder sb = new StringBuilder();
		int prevSize = 0;
		for (int i = 1; i <= 25000; i++) {
			Figure figure = dropFigure(Figure.getNextFigure(i, board.getYBoarder()), board, pattern);
			board.addFigure(figure);
			int currentSize = board.getCurrentSize();
			sizeAdded.put((long) i, (long) currentSize);
			int diff = currentSize - prevSize;
			sb.append(diff);
			prevSize = currentSize;
		}
		System.out.println("Silver: " + sizeAdded.get(2022L));
		String fullString = sb.toString();
		String repeat = longestRepeatingString(fullString);
		long repeatLength = repeat.length();
		long repeatStart = fullString.indexOf(repeat);
		long repeatSize = sizeAdded.get(repeatLength + repeatStart) - sizeAdded.get(repeatStart);

		long iteration = 1000000000000L;
		long answer =((iteration - repeatStart) / repeatLength) * repeatSize  + sizeAdded.get(repeatStart + (iteration - repeatStart) % repeatLength);;
		System.out.println("Gold: " + answer);
	}

	public Figure dropFigure(Figure figure, Board board, String pattern) {
		figure.pushGas(pattern, board);
		boolean canFailDown = figure.failDown(board);
		if (canFailDown) {
			return dropFigure(figure, board, pattern);
		} else {
			return figure;
		}
	}

	public void drawBoard(Board board) {
		int startX = 0;
		int endX = 7;
		int startY = 0;
		int endY = board.getYBoarder();

		System.out.println("=Draw= (" + startX + "->" + endX + ")  (" + startY + "->" + endY + ")");
		for (int y = endY; y >= startY; y--) {
			StringBuilder sb = new StringBuilder();
			for (int x = startX; x <= endX; x++) {
				String cur = board.points.contains(Point.of(x, y)) ? "#" : ".";
				sb.append(cur);
			}
			System.out.println(sb);
		}
	}

	public String longestRepeatingString(String str) {
		int n = str.length();
		int[][] LCSRe = new int[n + 1][n + 1];

		String res = "";
		int res_length = 0;
		int i, index = 0;
		for (i = 1; i <= n; i++) {
			for (int j = i + 1; j <= n; j++) {
				if (str.charAt(i - 1) == str.charAt(j - 1) && LCSRe[i - 1][j - 1] < (j - i)) {
					LCSRe[i][j] = LCSRe[i - 1][j - 1] + 1;
					if (LCSRe[i][j] > res_length) {
						res_length = LCSRe[i][j];
						index = Math.max(i, index);
					}
				} else {
					LCSRe[i][j] = 0;
				}
			}
		}
		if (res_length > 0) {
			for (i = index - res_length + 1; i <= index; i++) {
				res += str.charAt(i - 1);
			}
		}

		return res;
	}

	@Data static class Point {

		int x;
		int y;

		public static Point of(int x, int y) {
			Point point = new Point();
			point.x = x;
			point.y = y;
			return point;
		}

		public static Point addPoints(Point p, Point b) {
			return Point.of(b.x + p.x, b.y + p.y);
		}

		public void add(Point p) {
			this.x = this.x + p.x;
			this.y = this.y + p.y;
		}

		@Override public String toString() {
			return "{" + x + "," + y + '}';
		}
	}

	@Data static class Board {

		List<Point> points = new ArrayList<>();

		public void addFigure(Figure figure) {
			this.points.addAll(figure.getPoints());
		}

		public int getCurrentSize() {
			int maxY = 0;
			for (Point point : points) {
				maxY = Math.max(maxY, point.y + 1);
			}
			return maxY;
		}

		public int getYBoarder() {
			int maxY = 0;
			for (Point point : points) {
				maxY = Math.max(maxY, point.y + 1);
			}
			return maxY + 2;
		}
	}

	@Data static class Figure {

		List<Point> points = new ArrayList<>();

		public static Figure getNextFigure(int count, int startY) {
			count = count % 5;
			if (count == 1) {
				return getFirst(startY);
			}
			if (count == 2) {
				return getSecond(startY);
			}
			if (count == 3) {
				return getThird(startY);
			}
			if (count == 4) {
				return getFourth(startY);
			}
			if (count == 0) {
				return getFifth(startY);
			}
			return null;
		}

		public static Figure getFirst(int startY) {
			Figure figure = new Figure();
			figure.points.add(Point.of(2, 1 + startY));
			figure.points.add(Point.of(3, 1 + startY));
			figure.points.add(Point.of(4, 1 + startY));
			figure.points.add(Point.of(5, 1 + startY));
			return figure;
		}

		public static Figure getSecond(int startY) {
			Figure figure = new Figure();
			figure.points.add(Point.of(3, 1 + startY));
			figure.points.add(Point.of(2, 2 + startY));
			figure.points.add(Point.of(3, 2 + startY));
			figure.points.add(Point.of(4, 2 + startY));
			figure.points.add(Point.of(3, 3 + startY));
			return figure;
		}

		public static Figure getThird(int startY) {
			Figure figure = new Figure();
			figure.points.add(Point.of(4, 1 + startY));
			figure.points.add(Point.of(4, 2 + startY));
			figure.points.add(Point.of(4, 3 + startY));
			figure.points.add(Point.of(3, 1 + startY));
			figure.points.add(Point.of(2, 1 + startY));
			return figure;
		}

		public static Figure getFourth(int startY) {
			Figure figure = new Figure();
			figure.points.add(Point.of(2, 1 + startY));
			figure.points.add(Point.of(2, 2 + startY));
			figure.points.add(Point.of(2, 3 + startY));
			figure.points.add(Point.of(2, 4 + startY));
			return figure;
		}

		public static Figure getFifth(int startY) {
			Figure figure = new Figure();
			figure.points.add(Point.of(2, 1 + startY));
			figure.points.add(Point.of(3, 1 + startY));
			figure.points.add(Point.of(2, 2 + startY));
			figure.points.add(Point.of(3, 2 + startY));
			return figure;
		}

		public List<Point> getPoints() {
			return points;//new ArrayList<>(points);
		}

		public void pushGas(String pattern, Board board) {
			boolean isLeft = '<' == pattern.charAt(cursor % pattern.length());
			cursor++;
			int dx = isLeft ? -1 : 1;
			Point shift = Point.of(dx, 0);

			//			System.out.println("Jet of gas pushes rock " + ( isLeft ? "left " : "right " )+ shift);
			for (Point point : points) {
				Point newPoint = Point.addPoints(point, shift);
				if (board.points.contains(newPoint)) {
					return;
				}
				if (newPoint.x >= X_SIZE || newPoint.x < 0) {
					return;
				}
			}
			for (Point point : points) {
				point.add(shift);
			}
		}

		public boolean failDown(Board board) {
			Point shift = Point.of(0, -1);
			for (Point point : points) {

				Point newPoint = Point.addPoints(point, shift);
				if (newPoint.y < 0) {
					return false;
				}
				if (board.points.contains(newPoint)) {
					return false;
				}
			}
			//			System.out.println("Rock falls 1 unit:");
			for (Point point : points) {
				point.add(shift);
			}
			return true;
		}

	}



}
